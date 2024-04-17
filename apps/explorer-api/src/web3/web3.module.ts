import { Module } from '@nestjs/common';
import { Web3Service } from './web3.service';
import { BlocksModule } from 'src/blocks/blocks.module';
import { TxsModule } from 'src/txs/txs.module';
import { Web3Controller } from './web3.controller';

@Module({
  imports: [BlocksModule, TxsModule],
  controllers: [Web3Controller],
  providers: [Web3Service],
  exports: [Web3Service],
})
export class Web3Module {}
