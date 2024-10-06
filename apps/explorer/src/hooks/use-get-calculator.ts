import useSWR from 'swr';
import { carbonMonsterTestClient } from '@/utils/api-client';
import { CHAIN_ID_EMISSION } from '@/constants/chain';

export function useGetCalculator(params: { chainID: number; voucherID: string; tokenNo: string }) {
	const shouldFetch = params.chainID === CHAIN_ID_EMISSION;

	const { data, error, mutate } = useSWR(shouldFetch ? [`/footprints/${params.voucherID}/${params.tokenNo}`] : null, async (args) => (await carbonMonsterTestClient.get(args[0])).data);

	const resultData = data ? data.data : data;

	return { data: resultData, error, mutate };
}
