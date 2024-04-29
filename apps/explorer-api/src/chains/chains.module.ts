import { Module } from '@nestjs/common';
import { ChainsController } from './chains.controller';
import { ChainsService } from './chains.service';
import { Web3Module } from 'src/web3/web3.module';
import { HttpModule } from '@nestjs/axios';
import { VouchersModule } from 'src/vouchers/vouchers.module';

@Module({
  imports: [Web3Module, HttpModule, VouchersModule],
  controllers: [ChainsController],
  providers: [ChainsService],
  exports: [ChainsService],
})
export class ChainsModule {}
