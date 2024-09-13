import useSWR from 'swr';
import { carbonMonsterClient } from '@/utils/api-client';
import { CHAIN_ID_OFFSET } from '@/constants/chain';

export function useInfoTokens({ type, chainID }: { type: 'ft' | 'nft'; chainID: number }) {
	const shouldFetch = !(type === 'nft' && chainID !== CHAIN_ID_OFFSET);

	const { data, error, mutate } = useSWR(shouldFetch ? ['/projects/contracts'] : null, async (args) => (await carbonMonsterClient.get(args[0])).data);

	const resultData = data ? data.data : data;

	return { data: resultData, error, mutate };
}
