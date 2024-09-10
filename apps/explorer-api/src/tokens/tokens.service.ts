import { Injectable } from '@nestjs/common';
import { BalanceTokenRequestQueryDto, ListTokensRequestQueryDto, ListTokensResponseDto } from './dtos/list-tokens.dto';
import { Web3Service } from 'src/web3/web3.service';
import Web3 from 'web3';
import { TokenDto } from './dtos/token.dto';

@Injectable()
export class TokensService {
    constructor(private web3Service: Web3Service) {}

    async getFTData(chainID: number, address: string) {
        const provider = this.web3Service.getProvider(chainID);
        const token = new TokenDto();

        token.type = 'ERC20';
        token.deposit = '';
        token.address = address;

        const symbolABI = await provider.eth.call({
            to: address,
            data: provider.eth.abi.encodeFunctionCall(
                {
                    name: 'symbol',
                    type: 'function',
                    inputs: [],
                },
                [],
            ),
        });
        const symbol = provider.eth.abi.decodeParameter('string', symbolABI) as string;

        token.symbol = symbol;

        const totalSupplyABI = await provider.eth.call({
            to: address,
            data: provider.eth.abi.encodeFunctionCall(
                {
                    name: 'symbol',
                    type: 'function',
                    inputs: [],
                },
                [],
            ),
        });

        const totalSupply = BigInt(provider.eth.abi.decodeParameter('uint256', totalSupplyABI) as any);

        token.totalSupply = totalSupply.toString();

        const transferLogs = await provider.eth.getPastLogs({
            address: address,
            topics: [Web3.utils.sha3('Transfer(address,address,uint256)')],
            fromBlock: 0,
            toBlock: 'latest',
        });

        const holders = transferLogs.reduce((p: any, x: any) => {
            const from = provider.eth.abi.decodeParameter('address', x.topics[2]) as string;

            const to = provider.eth.abi.decodeParameter('address', x.topics[2]) as string;

            return Array.from(new Set([...p, ...[from, to]]));
        }, []);

        token.holders = holders.length;

        return token;
    }

    async getNFTData(chainID: number, address: string) {
        const provider = this.web3Service.getProvider(chainID);
        const token = new TokenDto();

        token.type = 'ERC1155';
        token.address = address;

        const voucherAddressABI = await provider.eth.call({
            to: address,
            data: provider.eth.abi.encodeFunctionCall(
                {
                    name: 'voucherAddress',
                    type: 'function',
                    inputs: [],
                },
                [],
            ),
        });
        const voucherAddress = provider.eth.abi.decodeParameter('address', voucherAddressABI) as string;

        token.deposit = voucherAddress;

        const symbolABI = await provider.eth.call({
            to: address,
            data: provider.eth.abi.encodeFunctionCall(
                {
                    name: 'symbol',
                    type: 'function',
                    inputs: [],
                },
                [],
            ),
        });
        const symbol = provider.eth.abi.decodeParameter('string', symbolABI) as string;

        token.symbol = symbol;

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

        token.totalSupply = (mint - burn).toString();

        const holders = logs.reduce((p: any, x: any) => {
            const from = provider.eth.abi.decodeParameter('address', x.topics[2]) as string;

            const to = provider.eth.abi.decodeParameter('address', x.topics[3]) as string;

            return Array.from(new Set([...p, ...[from, to]]));
        }, []);

        token.holders = holders.length;

        return token;
    }

    async listTokens(request: ListTokensRequestQueryDto): Promise<ListTokensResponseDto> {
        const func = request.type === 'ft' ? this.web3Service.getFTAddresses : this.web3Service.getNFTAddresses;

        const addresses = func(Number(request.chainID), Number(request.pageOffset), Number(request.pageSize));

        const results = await Promise.all(addresses.map((address) => (request.type === 'ft' ? this.getFTData(Number(request.chainID), address) : this.getNFTData(Number(request.chainID), address))));

        return {
            totalSize: results.length,
            tokens: results,
        };
    }

    async balanceTokens(request: BalanceTokenRequestQueryDto) {
        const results = await this.web3Service.balanceTokens(request.chainID, request.address);

        return results;
    }
}
