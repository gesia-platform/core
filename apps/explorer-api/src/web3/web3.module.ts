import { Module } from '@nestjs/common';
import { Web3Service } from './web3.service';
import { BlocksModule } from 'src/blocks/blocks.module';
import { TxsModule } from 'src/txs/txs.module';

@Module({
  imports: [BlocksModule, TxsModule],
  providers: [Web3Service],
  exports: [Web3Service],
})
export class Web3Module {}
