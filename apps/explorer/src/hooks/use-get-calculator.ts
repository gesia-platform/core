import useSWR from 'swr';
import { carbonMonsterTestClient } from '@/utils/api-client';
import { CHAIN_ID_EMISSION } from '@/constants/chain';

export function useGetCalculator(params: { chainID: number; emissionContract: string; pageOffset: number; pageSize: number }) {
	const shouldFetch = params.chainID === CHAIN_ID_EMISSION;

	const { data, error, mutate } = useSWR(shouldFetch ? [`/footprints/${params.emissionContract}?page=${params.pageOffset}&size=${params.pageSize}`] : null, async (args) => (await carbonMonsterTestClient.get(args[0])).data);

	const resultData = data ? data.data : data;

	console.log(resultData);

	return { data: resultData, error, mutate };
}
