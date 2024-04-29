import { Controller, Get, Param, Query } from '@nestjs/common';
import { ListVouchersRequestQueryDto } from './dtos/list-vouchers.dto';
import { GetVoucherRequestQueryDto } from './dtos/get-voucher.dto';
import { VouchersService } from './vouchers.service';
import { ListVoucherTokensRequestQueryDto } from './dtos/list-voucher-tokens.dto';

@Controller('/vouchers')
export class VouchersController {
  constructor(private vouchersService: VouchersService) {}

  @Get('')
  listVouchers(@Query() query: ListVouchersRequestQueryDto) {
    return this.vouchersService.listVouchers(query);
  }

  @Get('/:voucherID')
  getVoucher(
    @Query() query: GetVoucherRequestQueryDto,
    @Param('voucherID') voucherID: string,
  ) {
    return this.vouchersService.getVoucher(voucherID, query);
  }

  @Get('/:voucherID/tokens')
  listVoucherTokens(
    @Query() query: ListVoucherTokensRequestQueryDto,
    @Param('voucherID') voucherID: string,
  ) {
    return this.vouchersService.listVoucherTokens(voucherID, query);
  }
}
