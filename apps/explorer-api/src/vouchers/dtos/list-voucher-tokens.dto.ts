export class ListVoucherTokensRequestQueryDto {
    pageOffset: string;
    pageSize: string;
    chainID: string;
}

export class ListVoucherTokensResponseDto {
    totalSize: number;
    tokens: any[];
}
