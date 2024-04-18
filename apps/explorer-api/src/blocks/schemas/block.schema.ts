import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';

export type BlockDocument = HydratedDocument<Block>;

@Schema({ collection: 'blocks' })
export class Block {
  @Prop({ index: true })
  chainID: number;

  @Prop()
  height: string;

  @Prop()
  hash: string;

  @Prop()
  parentHash: string;

  @Prop()
  nonce: string;

  @Prop()
  timestamp: string;

  @Prop()
  miner: string;

  @Prop()
  difficulty: string;

  @Prop()
  gasUsed: string;

  @Prop()
  gasLimit: string;

  @Prop()
  baseFeePerGas: string;

  @Prop()
  extraData: string;

  @Prop()
  txns: number;

  @Prop()
  stateRoot: string;

  @Prop()
  transactionsRoot: string;

  @Prop()
  receiptsRoot: string;

  @Prop()
  logsBloom: string;

  @Prop()
  mixHash: string;

  @Prop()
  withdrawalsRoot: string;

  @Prop()
  size: string;

  @Prop()
  totalDifficulty: string;
}

export const BlockSchema = SchemaFactory.createForClass(Block);
