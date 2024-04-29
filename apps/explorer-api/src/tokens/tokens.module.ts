import { Module } from '@nestjs/common';
import { Web3Module } from 'src/web3/web3.module';
import { TokensController } from './tokens.controller';
import { TokensService } from './tokens.service';

@Module({
  imports: [Web3Module],
  controllers: [TokensController],
  providers: [TokensService],
})
export class TokensModule {}
