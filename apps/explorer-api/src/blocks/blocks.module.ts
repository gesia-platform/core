import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { Block, BlockSchema } from './schemas/block.schema';
import { BlocksService } from './blocks.service';
import { BlocksController } from './blocks.controller';
import { Web3Module } from 'src/web3/web3.module';

@Module({
    imports: [Web3Module, MongooseModule.forFeature([{ name: Block.name, schema: BlockSchema }])],
    controllers: [BlocksController],
    providers: [BlocksService],
    exports: [BlocksService],
})
export class BlocksModule {}
