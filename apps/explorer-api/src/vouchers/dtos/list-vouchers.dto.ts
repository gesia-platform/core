export class ListVouchersRequestQueryDto {
  pageOffset: string;
  pageSize: string;
  chainID: string;
}

export class ListVouchersResponseDto {
  totalSize: number;
  vouchers: any[];
}
