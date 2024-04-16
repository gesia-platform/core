import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { Tx, TxSchema } from './schemas/tx.schema';
import { TxsService } from './txs.service';
import { TxsController } from './txs.controller';

@Module({
  imports: [MongooseModule.forFeature([{ name: Tx.name, schema: TxSchema }])],
  controllers: [TxsController],
  providers: [TxsService],
  exports: [TxsService],
})
export class TxsModule {}
