import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import mongoose, { HydratedDocument } from 'mongoose';

export type TxDocument = HydratedDocument<Tx>;

@Schema({ collection: 'txs' })
export class Tx {
    @Prop({ type: mongoose.Schema.Types.ObjectId, ref: 'Block' })
    blockID: mongoose.Types.ObjectId;

    @Prop()
    hash: string;

    @Prop()
    from: string;

    @Prop()
    to: string;

    @Prop()
    contract: string;

    @Prop()
    gas: string;

    @Prop()
    gasPrice: string;

    @Prop()
    gasUsed: string;

    @Prop()
    maxFePerGas: string;

    @Prop()
    maxPriorityFeePerGas: string;

    @Prop()
    effectiveGasPrice: string;

    @Prop()
    cumulativeGasUsed: string;

    @Prop()
    input: string;

    @Prop()
    value: string;

    @Prop()
    nonce: string;

    @Prop()
    type: string;

    @Prop()
    v: string;
    @Prop()
    r: string;
    @Prop()
    s: string;

    @Prop()
    logs: any[];

    @Prop()
    logsBloom: string;

    @Prop()
    status: string;

    @Prop()
    transactionIndex: string;
}

export const TxSchema = SchemaFactory.createForClass(Tx);
