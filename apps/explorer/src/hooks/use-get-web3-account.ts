import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useGetWeb3Account(params?: {
  accountID: string;
  chainID: number;
}) {
  const { data, error, mutate } = useSWR(
    params ? ["/web3/accounts/" + params.accountID, params] : null,
    async (args) =>
      (
        await apiClient.get<{
          account: {
            address: string;
            balance: string;
            isContract: boolean;
            isIOA: boolean;
            isEmissionVoucher: boolean;
            isOffsetVoucher: boolean;
            isEmissionCalculator: boolean;
          };
        }>(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
