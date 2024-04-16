import { Module } from '@nestjs/common';
import { ChainsController } from './chains.controller';
import { ChainsService } from './chains.service';
import { Web3Module } from 'src/web3/web3.module';
import { HttpModule } from '@nestjs/axios';

@Module({
  imports: [Web3Module, HttpModule],
  controllers: [ChainsController],
  providers: [ChainsService],
  exports: [ChainsService],
})
export class ChainsModule {}
