import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Block } from './schemas/block.schema';
import mongoose, {
  Model,
  PipelineStage,
  isObjectIdOrHexString,
} from 'mongoose';
import {
  ListBlocksRequestQueryDto,
  ListBlocksResponseDto,
} from './dtos/list-blocks.dto';
import { GetBlockRequestQueryDto } from './dtos/get-block.dto';
import Web3 from 'web3';
import { Web3Service } from 'src/web3/web3.service';

@Injectable()
export class BlocksService {
  constructor(
    @InjectModel(Block.name) private blockModel: Model<Block>,
    private web3Service: Web3Service,
  ) {}

  async insertBlock(block: Block) {
    const blockModel = await this.blockModel.create(block);

    return await blockModel.save();
  }

  async exists(chainID: number, height: BigInt) {
    return await this.blockModel.exists({
      chainID,
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

  async getBlock(
    blockID: string,
    query: GetBlockRequestQueryDto,
    recursive?: boolean,
  ) {
    const pipelines: PipelineStage[] = [
      { $match: { chainID: Number(query.chainID) } },
      {
        $match: isObjectIdOrHexString(blockID)
          ? { _id: new mongoose.Types.ObjectId(blockID) }
          : { height: blockID },
      },
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
    ];

    const results = await this.blockModel.aggregate(pipelines);

    if (!results[0]) {
      if (!recursive) {
        await this.web3Service.processBlock(
          this.web3Service.getProvider(Number(query.chainID)),
          Number(query.chainID),
          blockID,
        );
        return await this.getBlock(blockID, query, true);
      }
      throw new NotFoundException();
    }

    return {
      block: results[0],
    };
  }

  async listBlocks(
    query: ListBlocksRequestQueryDto,
  ): Promise<ListBlocksResponseDto> {
    const provider = this.web3Service.getProvider(Number(query.chainID));

    const latestBlockNumber = Number(await provider.eth.getBlockNumber());

    const skip = Number(query.pageOffset);
    const limit = Number(query.pageSize);

    const results = await this.blockModel.aggregate([
      { $match: { chainID: Number(query.chainID) } },
      {
        $addFields: {
          heightLong: { $toLong: '$height' },
        },
      },
      { $sort: { heightLong: -1 } },
      { $skip: skip },
      { $limit: limit },
    ]);

    const blocks = [];

    const heights = [];

    for (
      let i = latestBlockNumber - skip;
      i > latestBlockNumber - skip - limit;
      i--
    ) {
      heights.push(i);
    }

    for await (let height of heights) {
      try {
        const block = (
          await this.getBlock(BigInt(height).toString(), {
            chainID: query.chainID,
          })
        ).block;
        blocks.push(block);
      } catch (error: unknown) {
        console.error(error);
      }
    }

    return {
      blocks: blocks,
      totalSize: latestBlockNumber,
    };
  }
}
