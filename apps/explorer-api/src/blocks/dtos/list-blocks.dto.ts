export class ListBlocksRequestQueryDto {
    pageOffset: string;
    pageSize: string;
    chainID: string;
}

export class ListBlocksResponseDto {
    blocks: any[];
    totalSize: number;
}
