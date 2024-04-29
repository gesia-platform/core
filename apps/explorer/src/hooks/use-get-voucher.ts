import { apiClient } from "@/utils/api-client";
import useSWR from "swr";

export function useGetVoucher(params: { voucherID: string; chainID: number }) {
  const { data, error, mutate } = useSWR(
    ["/vouchers/" + params.voucherID, params],
    async (args) =>
      (
        await apiClient.get(args[0], {
          params: args[1],
        })
      ).data
  );

  return { data, error, mutate };
}
