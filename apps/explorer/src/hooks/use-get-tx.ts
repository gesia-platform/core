import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useGetTx(params: { txID: string }) {
  const { data, error, mutate } = useSWR(
    ["/txs/" + params.txID, params],
    async (args) =>
      (
        await apiClient.get(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
