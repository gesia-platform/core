import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Tx } from './schemas/tx.schema';
import mongoose, { Model, PipelineStage, isObjectIdOrHexString } from 'mongoose';
import { ListTxsRequestQueryDto, ListTxsResponseDto } from './dtos/list-txs.dto';
import { GetTxRequestQueryDto } from './dtos/get-tx.dto';
import { Web3Service } from 'src/web3/web3.service';

@Injectable()
export class TxsService {
    constructor(@InjectModel(Tx.name) private txModel: Model<Tx>, private web3Service: Web3Service) {}

    async insertTx(tx: Tx) {
        const txModel = await this.txModel.create(tx);

        return await txModel.save();
    }

    async getTx(txID: string, query: GetTxRequestQueryDto, recursive?: boolean) {
        const pipelines: PipelineStage[] = [
            { $match: { hash: txID } },
            {
                $lookup: {
                    from: 'blocks',
                    let: {
                        blockID: '$blockID',
                    },
                    pipeline: [
                        {
                            $match: {
                                $expr: {
                                    $eq: ['$_id', '$$blockID'],
                                },
                            },
                        },
                        {
                            $project: {
                                timestamp: 1,
                                chainID: 1,
                                height: 1,
                            },
                        },
                    ],
                    as: 'block',
                },
            },
            {
                $unwind: {
                    path: '$block',
                    preserveNullAndEmptyArrays: true,
                },
            },
            {
                $addFields: {
                    timestamp: { $toInt: '$block.timestamp' },
                    index: { $toInt: '$transactionIndex' },
                },
            },
        ];

        const results = await this.txModel.aggregate([...pipelines]);

        const provider = this.web3Service.getProvider(Number(query.chainID));
        if (!results[0]) {
            if (!recursive) {
                await this.web3Service.processTransaction(provider, txID, Number(query.chainID));
                return await this.getTx(txID, query, true);
            }
            throw new NotFoundException();
        }

        const tx = results[0];

        if (!tx.to && !tx.contract) {
            const receipt = await this.web3Service.getTransactionReceipt(provider, txID);
            tx.contract = receipt.contractAddress;
            await this.txModel.updateOne({ _id: tx._id }, { contract: receipt.contractAddress });
        }

        return { tx };
    }

    async listTxs(query: ListTxsRequestQueryDto): Promise<ListTxsResponseDto> {
        const pipelines: PipelineStage[] = [
            {
                $lookup: {
                    from: 'blocks',
                    let: {
                        blockID: '$blockID',
                    },
                    pipeline: [
                        {
                            $match: {
                                $expr: {
                                    $eq: ['$_id', '$$blockID'],
                                },
                            },
                        },
                        {
                            $project: {
                                timestamp: 1,
                                chainID: 1,
                                height: 1,
                            },
                        },
                    ],
                    as: 'block',
                },
            },
            {
                $unwind: {
                    path: '$block',
                    preserveNullAndEmptyArrays: true,
                },
            },
            {
                $addFields: {
                    timestamp: { $toInt: '$block.timestamp' },
                    index: { $toInt: '$transactionIndex' },
                },
            },
            {
                $sort: {
                    timestamp: -1,
                    index: 1,
                },
            },
        ];

        if (query.chainID) {
            pipelines.push({
                $match: {
                    'block.chainID': Number(query.chainID),
                },
            });
        }

        if (query.blockID) {
            pipelines.push({
                $match: isObjectIdOrHexString(query.blockID) ? { blockID: new mongoose.Types.ObjectId(query.blockID) } : { 'block.height': query.blockID },
            });
        }

        const results = await this.txModel.aggregate([
            ...pipelines,
            {
                $facet: {
                    data: [{ $skip: Number(query.pageOffset) }, { $limit: Number(query.pageSize) }],
                    totalSize: [{ $count: 'count' }],
                },
            },
        ]);

        let txs = results[0]?.data ?? [];

        if (txs.filter((x) => !x.to && !x.contract).length >= 1) {
            const provider = this.web3Service.getProvider(Number(query.chainID));
            txs = await Promise.all(
                txs.map(async (tx) => {
                    try {
                        if (!tx.to && !tx.contract) {
                            const receipt = await this.web3Service.getTransactionReceipt(provider, tx.hash);
                            tx.contract = receipt.contractAddress;
                            await this.txModel.updateOne({ _id: tx._id }, { contract: receipt.contractAddress });
                        }
                    } catch (error: unknown) {
                        console.error('failed to update contract address for creation tx in list', error);
                    }

                    return tx;
                }),
            );
        }

        return {
            txs,
            totalSize: results[0]?.totalSize[0]?.count ?? 0,
        };
    }
}
