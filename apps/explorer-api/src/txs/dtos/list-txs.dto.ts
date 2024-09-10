export class ListTxsRequestQueryDto {
    pageOffset: string;
    pageSize: string;
    blockID: string;
    chainID: string;
}

export class ListTxsResponseDto {
    txs: any[];
    totalSize: number[];
}
