import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { BlocksModule } from './blocks/blocks.module';
import { ConfigModule } from '@nestjs/config';
import { Web3Module } from './web3/web3.module';
import { ChainsModule } from './chains/chains.module';
import { TxsModule } from './txs/txs.module';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MongooseModule.forRoot(process.env.MONGODB_URI),
    BlocksModule,
    TxsModule,
    Web3Module,
    ChainsModule,
  ],
})
export class AppModule {}
