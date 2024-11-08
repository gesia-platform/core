import { ConsoleLogger, Injectable, Logger } from '@nestjs/common';
import { Model } from 'mongoose';
import { Block } from 'src/blocks/schemas/block.schema';
import { Tx } from 'src/txs/schemas/tx.schema';
import Web3 from 'web3';
import { GetAccountRequestQueryDto } from './dtos/get-account.dto';
import { InjectModel } from '@nestjs/mongoose';
import { HttpService } from '@nestjs/axios';

@Injectable()
export class Web3Service {
    private readonly neutrality: Web3;
    private readonly emission: Web3;
    private readonly offset: Web3;
    private readonly logger = new Logger(Web3Service.name);

    private readonly emissionVoucherAddresses: string[];
    private readonly offsetVoucherAddresses: string[];

    private readonly emissionFTAddresses: string[];
    private readonly offsetFTAddresses: string[];

    private readonly emissionNFTAddresses: string[];
    private readonly offsetNFTAddresses: string[];

    private providers: { provider: Web3; chainID: number }[];

    constructor(@InjectModel(Tx.name) private txModel: Model<Tx>, @InjectModel(Block.name) private blockModel: Model<Block>, private httpService: HttpService) {
        this.neutrality = new Web3(process.env.CHAIN_NEUTRALITY_WS_URL);
        this.emission = new Web3(process.env.CHAIN_EMISSION_WS_URL);
        this.offset = new Web3(process.env.CHAIN_OFFSET_WS_URL);

        this.emissionVoucherAddresses = (process.env.CHAIN_EMISSION_VOUCHER_ADDRESSES || '').match(/.{1,42}/g) || [];
        this.offsetVoucherAddresses = (process.env.CHAIN_OFFSET_VOUCHER_ADDRESSES || '').match(/.{1,42}/g) || [];

        this.emissionFTAddresses = (process.env.CHAIN_EMISSION_FT_ADDRESSES || '').match(/.{1,42}/g) || [];
        this.offsetFTAddresses = (process.env.CHAIN_OFFSET_FT_ADDRESSES || '').match(/.{1,42}/g) || [];

        this.emissionNFTAddresses = (process.env.CHAIN_EMISSION_NFT_ADDRESSES || '').match(/.{1,42}/g) || [];
        this.offsetNFTAddresses = (process.env.CHAIN_OFFSET_NFT_ADDRESSES || '').match(/.{1,42}/g) || [];

        this.providers = [
            {
                provider: this.neutrality,
                chainID: 1,
            },
            {
                provider: this.emission,
                chainID: 2,
            },
            {
                provider: this.offset,
                chainID: 3,
            },
        ];

        if (process.env.WEB3_SUBSCRIBING === '1') {
            this.subscribeNewBlockHeaders().catch((err) => this.logger.error('failed to subscribeNewBlockHeaders', err));
        }

        if (process.env.WEB3_SYNCING === '1') {
            this.syncPreviousBlocks();
        }
    }

    getVoucherAddresses = (chainID: number, pageOffset?: number, pageSize?: number) => {
        const results = chainID === 2 ? this.emissionVoucherAddresses : chainID === 3 ? this.offsetVoucherAddresses : [];

        return pageOffset !== undefined ? results.slice(pageOffset, pageOffset + pageSize) : results;
    };

    getFTAddresses = (chainID: number, pageOffset?: number, pageSize?: number) => {
        const results = chainID === 2 ? this.emissionFTAddresses : chainID === 3 ? this.offsetFTAddresses : [];

        return pageOffset !== undefined ? results.slice(pageOffset, pageOffset + pageSize) : results;
    };

    getNFTAddresses = (chainID: number, pageOffset?: number, pageSize?: number) => {
        const results = chainID === 2 ? this.emissionNFTAddresses : chainID === 3 ? this.offsetNFTAddresses : [];

        return pageOffset !== undefined ? results.slice(pageOffset, pageOffset + pageSize) : results;
    };

    async subscribeNewBlockHeaders() {
        this.providers.forEach(async ({ provider, chainID }) => {
            const subscription = await provider.eth.subscribe('newBlockHeaders');

            subscription.on('error', (err) => {
                this.logger.error(`failed to subscribe chain #${chainID} newBlockHeaders`, err);
            });

            subscription.on('data', async (blockHeader) => {
                this.logger.log(`chain #${chainID} newBlockHeader #${blockHeader.number} (${blockHeader.hash})`);

                try {
                    await this.processBlock(provider, chainID, blockHeader.number.toString(), blockHeader.withdrawalsRoot);
                    this.logger.log(`Processed: chain #${chainID} block #${blockHeader.number.toString()}`);
                } catch (error: unknown) {
                    this.logger.error(`Error: chain #${chainID} block #${blockHeader.number.toString()}`, error);
                }
            });
        });
    }

    async syncPreviousBlocks() {
        this.providers.forEach(async ({ chainID, provider }) => {
            const chainLatestBlock = await provider.eth.getBlock();
            let i = BigInt(chainLatestBlock.number);
            while (i--) {
                if (i === BigInt(0)) {
                    continue;
                }

                const blockExists = await this.blockModel.exists({
                    chainID,
                    height: i.toString(),
                });

                if (!blockExists) {
                    try {
                        await this.processBlock(provider, chainID, i.toString());
                        this.logger.log(`Processed: chain #${chainID} block #${i.toString()}`);
                    } catch (error: unknown) {
                        this.logger.error(`Error: chain #${chainID} block #${i.toString()}`, error);
                    }
                } else {
                    this.logger.log(`Exists: chain #${chainID} block #${i.toString()}`);
                }
            }
        });
    }

    async saveBlock(provider: Web3, chainID: number, blockHeight: string, withdrawalsRoot?: string) {
        const blockData = await provider.eth.getBlock(blockHeight);

        let block = new Block();

        block.chainID = chainID;

        block.height = blockData.number.toString();
        block.height_int = Number(blockData.number);

        block.hash = blockData.hash;
        block.parentHash = blockData.parentHash;
        block.nonce = blockData.nonce.toString();
        block.timestamp = blockData.timestamp.toString();
        block.miner = blockData.miner;

        if (block.miner === '0x0000000000000000000000000000000000000000') {
            block.miner = await provider.eth.requestManager.send({
                method: 'clique_getSigner',
                params: [block.hash],
            });
        }
        block.difficulty = blockData.difficulty?.toString();
        block.gasUsed = blockData.gasUsed.toString();
        block.gasLimit = blockData.gasLimit.toString();
        block.baseFeePerGas = blockData.baseFeePerGas?.toString();
        block.extraData = blockData.extraData.toString();

        block.stateRoot = blockData.stateRoot;
        block.transactionsRoot = blockData.transactionsRoot;
        block.receiptsRoot = blockData.receiptsRoot;
        block.logsBloom = blockData.logsBloom?.toString();
        block.withdrawalsRoot = withdrawalsRoot;

        block.size = blockData.size?.toString();
        block.totalDifficulty = blockData.totalDifficulty?.toString();
        block.txns = blockData.transactions?.length ?? 0;

        block = await this.blockModel.create(block);

        return [block, blockData.transactions];
    }

    async processBlock(provider: Web3, chainID: number, blockHeight: string, withdrawalsRoot?: string) {
        const [block, transactions] = await this.saveBlock(provider, chainID, blockHeight, withdrawalsRoot);

        await Promise.all(
            (transactions as any)?.map(async (hash) => {
                return await this.processTransaction(provider, hash, chainID);
            }) ?? [],
        );
    }

    async getTransactionReceipt(provider: Web3, hash: string) {
        return await provider.eth.getTransactionReceipt(hash);
    }

    async processTransaction(provider: Web3, hash: string, chainID: number) {
        const transaction = await provider.eth.getTransaction(hash);
        const transactionReceipt = await provider.eth.getTransactionReceipt(hash);

        const tx = new Tx();

        let curBlock = await this.blockModel.findOne({
            height: transaction.blockNumber,
            chainID,
        });

        if (!curBlock) {
            curBlock = (await this.saveBlock(provider, chainID, transaction.blockNumber))[0] as any;
        }
        tx.blockID = curBlock._id;

        tx.hash = transaction.hash;
        tx.nonce = transaction.nonce;
        tx.transactionIndex = transaction.transactionIndex;
        tx.type = transaction.type;

        tx.value = transaction.value;
        tx.from = transaction.from;
        tx.to = transaction.to;
        tx.contract = transactionReceipt.contractAddress;
        tx.input = transaction.input;

        tx.gas = transaction.gas;
        tx.gasPrice = transaction.gasPrice;
        tx.gasUsed = transactionReceipt.gasUsed?.toString();
        tx.cumulativeGasUsed = transactionReceipt.cumulativeGasUsed?.toString();
        tx.effectiveGasPrice = transactionReceipt.effectiveGasPrice?.toString();
        tx.maxFePerGas = transaction.maxFeePerGas;
        tx.maxPriorityFeePerGas = transaction.maxPriorityFeePerGas;

        tx.logs = transactionReceipt.logs;
        tx.logsBloom = transactionReceipt.logsBloom?.toString();

        tx.r = transaction.r;
        tx.s = transaction.s;
        tx.v = transaction.v;

        tx.status = transactionReceipt.status?.toString();

        const txModel = await this.txModel.create(tx);

        await txModel.save();
    }

    async getAccount(accountID: string, query: GetAccountRequestQueryDto) {
        const provider = query.chainID === '1' ? this.getNeutrality() : query.chainID === '2' ? this.getEmission() : this.getOffset();

        const balance = await provider.eth.getBalance(accountID, undefined);

        const code = await provider.eth.getCode(accountID);
        const externalOwned = code == '0x';

        let isIOA: boolean;
        if (externalOwned) {
            let ioas = [];
            try {
                if (query.chainID === '2') {
                    const res = await this.httpService.axiosRef.get(process.env.CHAIN_EMISSION_GESIAD_URL + '/ioas');
                    ioas = res.data;
                } else if (query.chainID === '3') {
                    const res = await this.httpService.axiosRef.get(process.env.CHAIN_OFFSET_GESIAD_URL + '/ioas');
                    ioas = res.data;
                }

                if (ioas.includes(accountID)) {
                    isIOA = true;
                }
            } catch (error: any) {
                console.error(error.message);
            }
        }

        return {
            account: {
                address: accountID,
                balance: balance.toString(),
                isContract: !externalOwned,
                isIOA: isIOA,
            },
        };
    }

    getNeutrality(): Web3 {
        return this.neutrality;
    }

    getEmission(): Web3 {
        return this.emission;
    }

    getOffset(): Web3 {
        return this.offset;
    }

    getProvider(chainId: number): Web3 {
        switch (chainId) {
            case 1:
                return this.getNeutrality();
            case 2:
                return this.getEmission();
            case 3:
                return this.getOffset();
        }
    }

    async balanceTokens(chainID: string, address: string) {
        const ERC1155_ABI = [
            {
                inputs: [
                    {
                        internalType: 'address',
                        name: 'account',
                        type: 'address',
                    },
                    {
                        internalType: 'uint256',
                        name: 'id',
                        type: 'uint256',
                    },
                ],
                name: 'balanceOf',
                outputs: [
                    {
                        internalType: 'uint256',
                        name: '',
                        type: 'uint256',
                    },
                ],
                stateMutability: 'view',
                type: 'function',
            },
        ];

        const ERC1155_CONTRACTS = ['0x4273e90e9aCc5e1a6FB5a0D38dF3cdB1971C832F', '0x1121b15F9f4Da4F50086c204Ff248e35e582E9B7', '0x2896F27D13137DFD60e11867819492dFcA9Fb990'];

        const results: { contract: string; tokenId: number; balance: number }[] = [];
        const web3Provider = this.getProvider(Number(chainID));

        try {
            for (const contractAddress of ERC1155_CONTRACTS) {
                const nftContract = new web3Provider.eth.Contract(ERC1155_ABI, contractAddress);

                const balance = await nftContract.methods.balanceOf(address, 1).call();
                results.push({
                    contract: contractAddress,
                    tokenId: 1,
                    balance: Number(balance),
                });
            }
        } catch (error) {
            console.error('ERC-1155 NFT 잔액 조회 중 오류 발생:', error);
        }

        return results;
    }
}
