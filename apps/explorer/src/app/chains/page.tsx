import { Main } from "@/components/main";
import { ChainList } from "@/containers/chains/chain-list";

export default function Chains() {
  return (
    <Main>
      <h2 className="text-[22px] font-medium text-[#1C1E20]">Chains</h2>
      <ChainList />
    </Main>
  );
}
