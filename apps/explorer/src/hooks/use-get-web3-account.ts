import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useGetWeb3Account(params: {
  accountID: string;
  chainID: number;
}) {
  const { data, error, mutate } = useSWR(
    ["/web3/accounts/" + params.accountID, params],
    async (args) =>
      (
        await apiClient.get(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
