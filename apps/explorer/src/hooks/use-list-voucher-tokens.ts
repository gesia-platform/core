import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useListVoucherTokens(params?: {
  chainID: number;
  voucherID: string;
}) {
  const { data, error, mutate } = useSWR(
    params ? [`/vouchers/${params.voucherID}/tokens`, params] : null,
    async (args) =>
      (
        await apiClient.get(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
