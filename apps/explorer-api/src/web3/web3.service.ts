import { Injectable, Logger } from '@nestjs/common';
import mongoose from 'mongoose';
import { BlocksService } from 'src/blocks/blocks.service';
import { Block } from 'src/blocks/schemas/block.schema';
import { Tx } from 'src/txs/schemas/tx.schema';
import { TxsService } from 'src/txs/txs.service';
import Web3 from 'web3';

@Injectable()
export class Web3Service {
  private readonly neutrality: Web3;
  private readonly emission: Web3;
  private readonly offset: Web3;
  private readonly logger = new Logger(Web3Service.name);

  private providers: { provider: Web3; chainID: number }[];

  constructor(
    private blocksService: BlocksService,
    private txsService: TxsService,
  ) {
    this.neutrality = new Web3(process.env.CHAIN_NEUTRALITY_WS_URL);
    this.emission = new Web3(process.env.CHAIN_EMISSION_WS_URL);
    this.offset = new Web3(process.env.CHAIN_OFFSET_WS_URL);

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
      this.subscribeNewBlockHeaders().catch((err) =>
        this.logger.error('failed to subscribeNewBlockHeaders', err),
      );
    }

    if (process.env.WEB3_SYNCING === '1') {
      this.syncPreviousBlocks();
    }
  }

  async subscribeNewBlockHeaders() {
    this.providers.forEach(async ({ provider, chainID }) => {
      const subscription = await provider.eth.subscribe('newBlockHeaders');

      subscription.on('error', (err) => {
        this.logger.error(
          `failed to subscribe chain #${chainID} newBlockHeaders`,
          err,
        );
      });

      subscription.on('data', async (blockHeader) => {
        this.logger.log(
          `chain #${chainID} newBlockHeader #${blockHeader.number} (${blockHeader.hash})`,
        );

        try {
          await this.processBlock(
            provider,
            chainID,
            blockHeader.number.toString(),
            blockHeader.withdrawalsRoot,
          );
          this.logger.log(
            `chain #${chainID} block #${blockHeader.number.toString()} processed!`,
          );
        } catch (error: unknown) {
          this.logger.error(
            `chain #${chainID} block #${blockHeader.number.toString()} error`,
            error,
          );
        }
      });
    });
  }

  async syncPreviousBlocks() {
    this.providers.forEach(async ({ chainID, provider }) => {
      const chainLatestBlock = await provider.eth.getBlock();
      const originChainID = await provider.eth.getChainId();

      for (let i = BigInt(1); i < chainLatestBlock.number; i++) {
        if (!(await this.blocksService.exists(chainID, i))) {
          try {
            await this.processBlock(
              provider,
              chainID,
              originChainID.toString(),
              i.toString(),
            );
            this.logger.log(
              `chain #${chainID} block #${i.toString()} processed!`,
            );
          } catch (error: unknown) {
            this.logger.error(
              `chain #${chainID} block #${i.toString()} error`,
              error,
            );
          }
        }
      }
    });
  }

  async processBlock(
    provider: Web3,
    chainID: number,
    blockHeight: string,
    withdrawalsRoot?: string,
  ) {
    const blockData = await provider.eth.getBlock(blockHeight);

    const block = new Block();

    block.chainID = chainID;

    block.height = blockData.number.toString();
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

    const savedBlock = await this.blocksService.insertBlock(block);

    await Promise.all(
      blockData.transactions?.map(async (hash) => {
        return await this.processTransaction(provider, hash, savedBlock._id);
      }) ?? [],
    );
  }

  async processTransaction(
    provider: Web3,
    hash: string,
    blockID: mongoose.Types.ObjectId,
  ) {
    const transaction = await provider.eth.getTransaction(hash);
    const transactionReceipt = await provider.eth.getTransactionReceipt(hash);

    const tx = new Tx();

    tx.blockID = blockID;

    tx.hash = transaction.hash;
    tx.nonce = transaction.nonce;
    tx.transactionIndex = transaction.transactionIndex;
    tx.type = transaction.type;

    tx.value = transaction.value;
    tx.from = transaction.from;
    tx.to = transaction.to;
    tx.data = transaction.data;
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

    await this.txsService.insertTx(tx);
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
}
