import { apiClient } from '@/utils/api-client';
import useSWR from 'swr';

export function useListTokens(params: { chainID: number; type: 'ft' | 'nft'; pageOffset: number; pageSize: number }) {
	const { data, error, mutate } = useSWR(
		['/tokens', params],
		async (args) =>
			(
				await apiClient.get(args[0], {
					params: args[1],
				})
			).data,
	);

	return { data, error, mutate };
}
