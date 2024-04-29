import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useListVouchers(params: {
  chainID: number;
  pageOffset: number;
  pageSize: number;
}) {
  const { data, error, mutate } = useSWR(
    ["/vouchers", params],
    async (args) =>
      (
        await apiClient.get(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
