import { Controller, Get, Param, Query } from '@nestjs/common';
import { GetAccountRequestQueryDto } from './dtos/get-account.dto';
import { Web3Service } from './web3.service';

@Controller('/web3')
export class Web3Controller {
    constructor(private web3Service: Web3Service) {}
    @Get('/accounts/:accountID')
    getAccount(@Param('accountID') accountID: string, @Query() query: GetAccountRequestQueryDto) {
        return this.web3Service.getAccount(accountID, query);
    }
}
