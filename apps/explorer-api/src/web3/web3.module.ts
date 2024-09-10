import { Module } from '@nestjs/common';
import { Web3Service } from './web3.service';
import { Web3Controller } from './web3.controller';
import { MongooseModule } from '@nestjs/mongoose';
import { Block, BlockSchema } from 'src/blocks/schemas/block.schema';
import { Tx, TxSchema } from 'src/txs/schemas/tx.schema';
import { HttpModule } from '@nestjs/axios';

@Module({
    imports: [
        HttpModule,
        MongooseModule.forFeature([
            { name: Block.name, schema: BlockSchema },
            { name: Tx.name, schema: TxSchema },
        ]),
    ],
    controllers: [Web3Controller],
    providers: [Web3Service],
    exports: [Web3Service],
})
export class Web3Module {}
