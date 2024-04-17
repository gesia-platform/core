import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Tx } from './schemas/tx.schema';
import mongoose, {
  Model,
  PipelineStage,
  isObjectIdOrHexString,
} from 'mongoose';
import {
  ListTxsRequestQueryDto,
  ListTxsResponseDto,
} from './dtos/list-txs.dto';

@Injectable()
export class TxsService {
  constructor(@InjectModel(Tx.name) private txModel: Model<Tx>) {}

  async insertTx(tx: Tx) {
    const txModel = await this.txModel.create(tx);

    return await txModel.save();
  }

  async listTxs(query: ListTxsRequestQueryDto): Promise<ListTxsResponseDto> {
    const pipelines: PipelineStage[] = [
      {
        $lookup: {
          from: 'blocks',
          let: {
            blockID: '$blockID',
          },
          pipeline: [
            {
              $match: {
                $expr: {
                  $eq: ['$_id', '$$blockID'],
                },
              },
            },
            {
              $project: {
                timestamp: 1,
                chainID: 1,
                height: 1,
              },
            },
          ],
          as: 'block',
        },
      },
      {
        $unwind: {
          path: '$block',
          preserveNullAndEmptyArrays: true,
        },
      },
      {
        $addFields: {
          timestamp: { $toInt: '$block.timestamp' },
          index: { $toInt: '$transactionIndex' },
        },
      },
      {
        $sort: {
          timestamp: -1,
          index: 1,
        },
      },
    ];

    if (query.chainID) {
      pipelines.push({
        $match: {
          'block.chainID': Number(query.chainID),
        },
      });
    }

    if (query.blockID) {
      pipelines.push({
        $match: isObjectIdOrHexString(query.blockID)
          ? { blockID: new mongoose.Types.ObjectId(query.blockID) }
          : { 'block.height': query.blockID },
      });
    }

    const results = await this.txModel.aggregate([
      ...pipelines,
      {
        $facet: {
          data: [
            { $skip: Number(query.pageOffset) },
            { $limit: Number(query.pageSize) },
          ],
          totalSize: [{ $count: 'count' }],
        },
      },
    ]);

    return {
      txs: results[0]?.data ?? [],
      totalSize: results[0]?.totalSize[0]?.count ?? 0,
    };
  }
}
