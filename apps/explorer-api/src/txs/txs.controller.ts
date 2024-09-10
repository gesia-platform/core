import { Controller, Get, Param, Query } from '@nestjs/common';
import { TxsService } from './txs.service';
import { ListTxsRequestQueryDto } from './dtos/list-txs.dto';

@Controller('/txs')
export class TxsController {
    constructor(private txsService: TxsService) {}

    @Get('/:txID')
    getTx(@Query() query: ListTxsRequestQueryDto, @Param('txID') txID: string) {
        return this.txsService.getTx(txID, query);
    }

    @Get('')
    listTxs(@Query() query: ListTxsRequestQueryDto) {
        return this.txsService.listTxs(query);
    }
}
