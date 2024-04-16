import { Controller, Get, Query } from '@nestjs/common';
import { TxsService } from './txs.service';
import { ListTxsRequestQueryDto } from './dtos/list-txs.dto';

@Controller('/txs')
export class TxsController {
  constructor(private txsService: TxsService) {}

  @Get('')
  listTxs(@Query() query: ListTxsRequestQueryDto) {
    return this.txsService.listTxs(query);
  }
}
