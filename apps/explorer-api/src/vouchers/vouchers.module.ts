import { Module } from '@nestjs/common';
import { Web3Module } from 'src/web3/web3.module';
import { VouchersController } from './vouchers.controller';
import { VouchersService } from './vouchers.service';

@Module({
    imports: [Web3Module],
    controllers: [VouchersController],
    exports: [VouchersService],
    providers: [VouchersService],
})
export class VouchersModule {}
