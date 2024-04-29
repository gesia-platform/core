import { Main } from "@/components/main";
import { PageLabel } from "@/components/page-label";
import { VoucherList } from "@/containers/voucher/list";

export default function Vouchers() {
  return (
    <Main>
      <PageLabel>Carbon Voucher Tracker</PageLabel>
      <VoucherList />
    </Main>
  );
}
