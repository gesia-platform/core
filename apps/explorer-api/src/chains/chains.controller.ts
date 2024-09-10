import { Controller, Get } from '@nestjs/common';
import { ChainsService } from './chains.service';

@Controller('/chains')
export class ChainsController {
    constructor(private chainsService: ChainsService) {}

    @Get('')
    async listChains() {
        return await this.chainsService.listChains();
    }
}
