import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { BlocksModule } from './blocks/blocks.module';
import { ConfigModule } from '@nestjs/config';
import { Web3Module } from './web3/web3.module';
import { ChainsModule } from './chains/chains.module';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MongooseModule.forRoot(process.env.MONGODB_URI),
    BlocksModule,
    Web3Module,
    ChainsModule,
  ],
})
export class AppModule {}
