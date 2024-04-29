import { Injectable } from '@nestjs/common';
import { ChainsConstants } from './chains.constant';
import { Web3Service } from 'src/web3/web3.service';
import Web3 from 'web3';
import { HttpService } from '@nestjs/axios';

@Injectable()
export class ChainsService {
  constructor(
    private web3Service: Web3Service,
    private httpService: HttpService,
  ) {
    this.listChains();
  }

  async listChains() {
    const latestBlockHeightNeutrality = await this.web3Service
      .getNeutrality()
      .eth.getBlockNumber();

    const latestBlockHeightEmission = await this.web3Service
      .getEmission()
      .eth.getBlockNumber();

    const latestBlockHeightOffset = await this.web3Service
      .getOffset()
      .eth.getBlockNumber();

    const emissionSigners = await this.web3Service
      .getEmission()
      .eth.requestManager.send({
        method: 'clique_getSigners',
        params: [Web3.utils.toHex(latestBlockHeightEmission)],
      });

    const offsetSigners = await this.web3Service
      .getOffset()
      .eth.requestManager.send({
        method: 'clique_getSigners',
        params: [Web3.utils.toHex(latestBlockHeightOffset)],
      });

    this.httpService.axiosRef.defaults.baseURL = process.env
      .CHAIN_NEUTRALITY_BEACON_API_URL as string;

    const validatorsRes = await this.httpService.axiosRef.get(
      `/eth/v1/beacon/states/head/validators`,
    );

    return {
      totalSize: 3,
      chains: [
        {
          id: ChainsConstants.NEUTRALITY_ID,
          name: ChainsConstants.NEUTRALITY_NAME,
          label: ChainsConstants.NEUTRALITY_LABEL,
          consensusAlgorithm: 'PoS',
          latestBlockHeight: latestBlockHeightNeutrality.toString(),
          nodes: validatorsRes.data.data?.length ?? 0,
          carbonTotalAmount: 0,
        },
        {
          id: ChainsConstants.EMISSION_ID,
          name: ChainsConstants.EMISSION_NAME,
          label: ChainsConstants.EMISSION_LABEL,
          consensusAlgorithm: 'PoA',
          latestBlockHeight: latestBlockHeightEmission.toString(),
          nodes: emissionSigners?.length ?? 0,
          carbonTotalAmount: await this.web3Service.getVoucherTotalAmount(
            ChainsConstants.EMISSION_ID,
          ),
        },
        {
          id: ChainsConstants.OFFSET_ID,
          name: ChainsConstants.OFFSET_NAME,
          label: ChainsConstants.OFFSET_LABEL,
          consensusAlgorithm: 'PoA',
          latestBlockHeight: latestBlockHeightOffset.toString(),
          nodes: offsetSigners?.length ?? 0,
          carbonTotalAmount: await this.web3Service.getVoucherTotalAmount(
            ChainsConstants.OFFSET_ID,
          ),
        },
      ],
    };
  }
}
