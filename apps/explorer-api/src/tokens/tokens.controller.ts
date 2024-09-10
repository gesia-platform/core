import { Controller, Get, Query } from '@nestjs/common';
import { BalanceTokenRequestQueryDto, ListTokensRequestQueryDto } from './dtos/list-tokens.dto';
import { TokensService } from './tokens.service';

@Controller('/tokens')
export class TokensController {
    constructor(private tokensService: TokensService) {}

    @Get('')
    listTokens(@Query() query: ListTokensRequestQueryDto) {
        return this.tokensService.listTokens(query);
    }

    @Get('/balance')
    balanceTokens(@Query() query: BalanceTokenRequestQueryDto) {
        return this.tokensService.balanceTokens(query);
    }
}
