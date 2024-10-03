import { CHAIN_ID_OFFSET } from '@/constants/chain';
import { apiClient } from '@/utils/api-client';
import useSWR from 'swr';

export function useGetCredit(params: { creditID: string; chainID: number; tokenNo: string; pageOffset: number; pageSize: number }) {
	const { data, error, mutate } = useSWR([`/credits`, params], async (args) => {
		const response = await apiClient.get(args[0], { params: args[1] });
		return response.data;
	});

	return { data, error, mutate };
}
