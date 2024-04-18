import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { Tx, TxSchema } from './schemas/tx.schema';
import { TxsService } from './txs.service';
import { TxsController } from './txs.controller';
import { Web3Module } from 'src/web3/web3.module';

@Module({
  imports: [
    Web3Module,
    MongooseModule.forFeature([{ name: Tx.name, schema: TxSchema }]),
  ],
  controllers: [TxsController],
  providers: [TxsService],
  exports: [TxsService],
})
export class TxsModule {}
