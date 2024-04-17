import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useGetBlock(params: { blockID: string; chainID: number }) {
  const { data, error, mutate } = useSWR(
    ["/blocks/" + params.blockID, params],
    async (args) =>
      (
        await apiClient.get(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
