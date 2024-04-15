import { Main } from "@/components/main";
import { TxList } from "@/containers/tx/list";

export default function Txs() {
  return (
    <Main>
      <h2 className="text-[22px] font-medium text-[#1C1E20]">Transactions</h2>
      <TxList />
    </Main>
  );
}
