import { Controller, Get, Query } from '@nestjs/common';
import { ListBlocksRequestQueryDto } from './dtos/list-blocks.dto';
import { BlocksService } from './blocks.service';

@Controller('/blocks')
export class BlocksController {
  constructor(private blocksService: BlocksService) {}

  @Get('')
  listBlocks(@Query() query: ListBlocksRequestQueryDto) {
    return this.blocksService.listBlocks(query);
  }
}
