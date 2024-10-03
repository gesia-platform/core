import { Controller, Get, Query } from '@nestjs/common';
import { CreditService } from './credits.service';
import { GetCreditRequestQueryDto } from './dtos/get-credit.dto';

@Controller('/credits')
export class CreditsController {
    constructor(private creditService: CreditService) {}

    @Get('')
    getCredit(@Query() query: GetCreditRequestQueryDto) {
        return this.creditService.getCredit(query);
    }
}
