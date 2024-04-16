import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useListChains(params: {}) {
  const { data, error, mutate } = useSWR(
    ["/chains", params],
    async (args) =>
      (
        await apiClient.get(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
