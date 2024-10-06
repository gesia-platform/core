import { Module } from '@nestjs/common';
import { Web3Module } from 'src/web3/web3.module';
import { CalculatorsController } from './calculators.controller';
import { CalculatorsService } from './calculators.service';

@Module({
    imports: [Web3Module],
    controllers: [CalculatorsController],
    exports: [CalculatorsService],
    providers: [CalculatorsService],
})
export class CalculatorsModule {}
