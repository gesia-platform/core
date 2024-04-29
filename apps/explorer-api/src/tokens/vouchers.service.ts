import { Injectable } from '@nestjs/common';
import { ListVouchersRequestQueryDto } from './dtos/list-vouchers.dto';
import { GetVoucherRequestQueryDto } from './dtos/get-voucher.dto';
import { Web3Service } from 'src/web3/web3.service';

@Injectable()
export class VouchersService {
  constructor(private web3Service: Web3Service) {}

  listVouchers(request: ListVouchersRequestQueryDto) {}

  getVoucher(voucherID: string, request: GetVoucherRequestQueryDto) {}
}
