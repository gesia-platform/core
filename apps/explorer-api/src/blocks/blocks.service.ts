import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Block } from './schemas/block.schema';
import { Model } from 'mongoose';
import {
  ListBlocksRequestQueryDto,
  ListBlocksResponseDto,
} from './dtos/list-blocks.dto';

@Injectable()
export class BlocksService {
  constructor(@InjectModel(Block.name) private blockModel: Model<Block>) {}

  async insertBlock(block: Block) {
    const blockModel = await this.blockModel.create(block);

    return await blockModel.save();
  }

  async exists(networkID: number, height: BigInt) {
    return await this.blockModel.exists({
      networkID,
      height: height.toString(),
    });
  }

  async getLatestBlock(networkID: number) {
    const block = await this.blockModel.findOne(
      { networkID },
      {},
      { sort: { height: -1 } },
    );
    return block;
  }

  async listBlocks(
    query: ListBlocksRequestQueryDto,
  ): Promise<ListBlocksResponseDto> {
    const results = await this.blockModel.aggregate([
      { $match: { chainID: Number(query.chainID) } },
      {
        $lookup: {
          from: 'txs',
          localField: '_id',
          foreignField: 'blockID',
          as: 'txs',
        },
      },
      {
        $addFields: {
          txns: { $size: '$txs' },
          heightLong: { $toLong: '$height' },
        },
      },

      { $sort: { heightLong: -1 } },
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
      blocks: results[0]?.data ?? [],
      totalSize: results[0]?.totalSize[0]?.count ?? 0,
    };
  }
}
