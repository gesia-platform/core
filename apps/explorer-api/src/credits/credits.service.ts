import { Injectable } from '@nestjs/common';
import { Web3Service } from 'src/web3/web3.service';
import Web3 from 'web3';
import { GetCreditRequestQueryDto } from './dtos/get-credit.dto';

@Injectable()
export class CreditService {
    constructor(private web3Service: Web3Service) {}

    public async getCredit(query: GetCreditRequestQueryDto): Promise<TransferEvent[]> {
        if (Number(query.chainID) !== 3) return;

        const provider = this.web3Service.getProvider(Number(query.chainID));

        // const contract = new provider.eth.Contract(
        //     [
        //         {
        //             inputs: [],
        //             name: 'owner',
        //             outputs: [{ internalType: 'address', name: '', type: 'address' }],
        //             stateMutability: 'view',
        //             type: 'function',
        //         },
        //     ],
        //     param.creditID,
        // );

        // const owner = await contract.methods.owner().call();

        const logs = await provider.eth.getPastLogs({
            address: query.creditID,
            topics: [Web3.utils.sha3('TransferSingle(address,address,address,uint256,uint256)')],
            fromBlock: 0,
            toBlock: 'latest',
        });

        const results = (
            await Promise.all(
                logs.map(async (x: any) => {
                    const datas = x.data.replace('0x', '').match(/.{1,64}/g);
                    const block = await provider.eth.getBlock(x.blockNumber);

                    if (!block) {
                        throw new Error('블록 정보를 가져오는 데 실패했습니다.');
                    }

                    const amount = Number(provider.eth.abi.decodeParameter('uint256', datas[1]));

                    if (amount === 0) return null;

                    return {
                        transactionHash: x.transactionHash,
                        methodName: 'TransferSingle',
                        blockNumber: Number(x.blockNumber),
                        creationTime: Number(block.timestamp),
                        from: provider.eth.abi.decodeParameter('address', x.topics[2]),
                        to: provider.eth.abi.decodeParameter('address', x.topics[3]),
                        amount,
                        id: Number(provider.eth.abi.decodeParameter('uint256', datas[0])),
                    } as unknown as TransferEvent;
                }),
            )
        ).filter((x) => x !== null);

        results.sort((a, b) => b.creationTime - a.creationTime);

        return results;
    }
}
