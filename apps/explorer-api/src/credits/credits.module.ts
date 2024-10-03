import { Module } from '@nestjs/common';
import { Web3Module } from 'src/web3/web3.module';
import { CreditsController } from './credits.controller';
import { CreditService } from './credits.service';

@Module({
    imports: [Web3Module],
    controllers: [CreditsController],
    exports: [CreditService],
    providers: [CreditService],
})
export class CreditsModule {}
