import { Controller, Get, Param, Query } from '@nestjs/common';
import { ListVouchersRequestQueryDto } from './dtos/list-vouchers.dto';
import { GetVoucherRequestQueryDto } from './dtos/get-voucher.dto';
import { VouchersService } from './vouchers.service';

@Controller('/vouchers')
export class VouchersController {
  constructor(private vouchersService: VouchersService) {}

  @Get('')
  listVouchers(@Query() query: ListVouchersRequestQueryDto) {}

  @Get('/:voucherID')
  getVoucher(
    @Query() query: GetVoucherRequestQueryDto,
    @Param('voucherID') voucherID: string,
  ) {}
}
