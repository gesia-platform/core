import { Injectable } from '@nestjs/common';
import { ListVouchersRequestQueryDto, ListVouchersResponseDto } from './dtos/list-vouchers.dto';
import { GetVoucherRequestQueryDto } from './dtos/get-voucher.dto';
import { Web3Service } from 'src/web3/web3.service';
import { VoucherDto } from './dtos/voucher.dto';
import Web3 from 'web3';
import { ListVoucherTokensRequestQueryDto, ListVoucherTokensResponseDto } from './dtos/list-voucher-tokens.dto';

@Injectable()
export class VouchersService {
    constructor(private web3Service: Web3Service) {}

    async getVoucherData(chainID: number, address: string) {
        const provider = this.web3Service.getProvider(chainID);

        const voucher = new VoucherDto();
        voucher.address = address;
        voucher.type = 'ERC1155';

        // count
        const mintLogs = await provider.eth.getPastLogs({
            address: address,
            topics: [Web3.utils.sha3('TransferSingle(address,address,address,uint256,uint256)'), null, '0x0000000000000000000000000000000000000000000000000000000000000000'],
            fromBlock: 0,
            toBlock: 'latest',
        });

        const tokenIds = new Set(
            mintLogs.map((x: any) => {
                const datas = x.data.replace('0x', '').match(/.{1,64}/g);
                return BigInt(provider.eth.abi.decodeParameter('uint256', datas[0]) as any).toString();
            }),
        );

        const nameABI = await provider.eth.call({
            to: address,
            data: provider.eth.abi.encodeFunctionCall(
                {
                    name: chainID === 2 ? 'getName' : 'name',
                    type: 'function',
                    inputs: [],
                },
                [],
            ),
        });
        const name = provider.eth.abi.decodeParameter('string', nameABI) as string;

        voucher.count = tokenIds.size;
        voucher.name = name;

        return voucher;
    }

    async listVoucherTokens(voucherID: string, request: ListVoucherTokensRequestQueryDto): Promise<ListVoucherTokensResponseDto> {
        const provider = this.web3Service.getProvider(Number(request.chainID));

        const logs = await provider.eth.getPastLogs({
            address: voucherID,
            topics: [Web3.utils.sha3('TransferSingle(address,address,address,uint256,uint256)')],
            fromBlock: 0,
            toBlock: 'latest',
        });

        const results = logs.map((x: any) => {
            const datas = x.data.replace('0x', '').match(/.{1,64}/g);
            return {
                to: provider.eth.abi.decodeParameter('address', x.topics[3]) as string,
                amount: BigInt(provider.eth.abi.decodeParameter('uint256', datas[1]) as any),
                id: BigInt(provider.eth.abi.decodeParameter('uint256', datas[0]) as any).toString(),
            };
        });

        const data = results.reduce(
            (p, c) => ({
                ...p,
                [c.id]: {
                    ...p[c.id],
                    [c.to]: c.amount + ((p[c.id] || {})[c.to] || BigInt(0)),
                },
            }),
            {},
        );

        const tokens = Object.keys(data).reduce((p, id) => {
            return [
                ...p,
                ...Object.entries(data[id]).map(([k, v]) => ({
                    owner: k,
                    amount: v.toString(),
                    id,
                })),
            ];
        }, []);

        return { tokens, totalSize: tokens.length };
    }

    async listVouchers(request: ListVouchersRequestQueryDto): Promise<ListVouchersResponseDto> {
        const addresses = this.web3Service.getVoucherAddresses(Number(request.chainID), Number(request.pageOffset), Number(request.pageSize));

        const results = await Promise.all(addresses.map((address) => this.getVoucherData(Number(request.chainID), address)));

        return {
            totalSize: results.length,
            vouchers: results,
        };
    }

    async getVoucher(voucherID: string, request: GetVoucherRequestQueryDto) {
        return {
            voucher: await this.getVoucherData(Number(request.chainID), voucherID),
        };
    }

    async getVoucherTotalAmount(chainID: number) {
        const provider = this.web3Service.getProvider(chainID);

        const addresses = this.web3Service.getVoucherAddresses(chainID);

        try {
            const results = await Promise.all(
                addresses.map(async (address): Promise<bigint> => {
                    const logs = await provider.eth.getPastLogs({
                        address: address,
                        topics: [Web3.utils.sha3('TransferSingle(address,address,address,uint256,uint256)')],
                        fromBlock: 0,
                        toBlock: 'latest',
                    });

                    const mints = [];
                    const burns = [];

                    logs.forEach((log: any) => {
                        if (log.topics[2] === '0x0000000000000000000000000000000000000000000000000000000000000000') {
                            mints.push(log);
                        } else if (log.topics[3] === '0x0000000000000000000000000000000000000000000000000000000000000000') {
                            burns.push(log);
                        }
                    });

                    const mint: bigint = mints.reduce((p: bigint, c) => {
                        const datas = c.data.replace('0x', '').match(/.{1,64}/g);
                        const value = BigInt(provider.eth.abi.decodeParameter('uint256', datas[1]) as any);

                        return p + value;
                    }, BigInt(0));

                    const burn: bigint = burns.reduce((p: bigint, c) => {
                        const datas = c.data.replace('0x', '').match(/.{1,64}/g);
                        const value = BigInt(provider.eth.abi.decodeParameter('uint256', datas[1]) as any);

                        return p + value;
                    }, BigInt(0));

                    return mint - burn;
                }),
            );

            return results.reduce((p, c) => p + c, BigInt(0)).toString();
        } catch (error) {
            console.log('getVoucherTotalAmount Error', error);
        }
    }
}
