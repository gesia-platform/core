import useSWR from 'swr';
import { carbonMonsterClient } from '@/utils/api-client';
import { CHAIN_ID_OFFSET } from '@/constants/chain';

export function useInfoToken({ type, chainID, nft_smart_contract }: { type: 'ft' | 'nft'; chainID: number; nft_smart_contract: string }) {
	const shouldFetch = !(type === 'nft' && chainID !== CHAIN_ID_OFFSET);

	const { data, error, mutate } = useSWR(shouldFetch ? [`/nfts/${nft_smart_contract}`] : null, async (args) => (await carbonMonsterClient.get(args[0])).data);

	const resultData = data ? data.data : data;

	return { data: resultData, error, mutate };
}
