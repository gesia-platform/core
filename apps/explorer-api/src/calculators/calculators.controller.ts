import { Controller, Get, Query } from '@nestjs/common';
import { CalculatorsService } from './calculators.service';
import { GetCalculatorRequestQueryDto } from './dtos/get-calculator.dto';

@Controller('/calculators')
export class CalculatorsController {
    constructor(private calculatorService: CalculatorsService) {}

    @Get('')
    getCredit(@Query() query: GetCalculatorRequestQueryDto) {
        return this.calculatorService.getCalculators(query);
    }
}
