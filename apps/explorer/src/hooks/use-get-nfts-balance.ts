import useSWR from 'swr';
import { apiClient } from '@/utils/api-client';

export function useGetNftsBalance(params: { chainID: number; userAddress: string }) {
	const { data, error, mutate } = useSWR(params ? `/tokens/balance?chainID=${params.chainID}&address=${params.userAddress}` : null, async (url) => (await apiClient.get(url)).data);

	return { data, error, mutate };
}
