import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useListBlocks(params: {
  chainID: number;
  pageOffset: number;
  pageSize: number;
}) {
  const { data, error, mutate } = useSWR(
    ["/blocks", params],
    async (args) =>
      (
        await apiClient.get(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
