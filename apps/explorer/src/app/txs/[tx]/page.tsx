import { Main } from "@/components/main";
import { PageLabel } from "@/components/page-label";
import { TxDetails } from "@/containers/tx/details";

export default function TxDetail({ params }: { params: any }) {
  return (
    <Main>
      <PageLabel>Transaction Details</PageLabel>
      <TxDetails txID={params.tx} />
    </Main>
  );
}
