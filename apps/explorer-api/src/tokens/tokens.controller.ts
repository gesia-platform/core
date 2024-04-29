import { Controller, Get, Param, Query } from '@nestjs/common';
import { ListTokensRequestQueryDto } from './dtos/list-tokens.dto';
import { TokensService } from './tokens.service';

@Controller('/tokens')
export class TokensController {
  constructor(private tokensService: TokensService) {}

  @Get('')
  listTokens(@Query() query: ListTokensRequestQueryDto) {
    return this.tokensService.listTokens(query);
  }
}
