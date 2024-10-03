export class GetCreditRequestQueryDto {
    creditID: string;
    tokenNo: string;
    chainID: string;
    pageOffset: number;
    pageSize: number;
}

export class GetCreditResponseDto {
    credit: any;
}
