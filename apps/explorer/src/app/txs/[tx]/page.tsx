import { Main } from "@/components/main";
import { TxDetails } from "@/containers/tx/details";

export default function TxDetail({ params }: { params: any }) {
  return (
    <Main>
      <h2 className="text-[22px] font-medium text-[#1C1E20]">
        Transaction Details
      </h2>
      <TxDetails txID={params.tx} />
    </Main>
  );
}
