export class ListTokensRequestQueryDto {
    pageOffset: string;
    pageSize: string;
    chainID: string;
    type: 'nft' | 'ft';
}

export class ListTokensResponseDto {
    totalSize: number;
    tokens: any[];
}

export class BalanceTokenRequestQueryDto {
    chainID: string;
    address: string;
}
