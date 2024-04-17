import { Controller, Get, Param, Query } from '@nestjs/common';
import { ListBlocksRequestQueryDto } from './dtos/list-blocks.dto';
import { BlocksService } from './blocks.service';
import { GetBlockRequestQueryDto } from './dtos/get-block.dto';

@Controller('/blocks')
export class BlocksController {
  constructor(private blocksService: BlocksService) {}

  @Get(':blockID')
  getBlockByNumber(
    @Query() query: GetBlockRequestQueryDto,
    @Param('blockID') blockID: string,
  ) {
    return this.blocksService.getBlock(blockID, query);
  }

  @Get('')
  listBlocks(@Query() query: ListBlocksRequestQueryDto) {
    return this.blocksService.listBlocks(query);
  }
}
