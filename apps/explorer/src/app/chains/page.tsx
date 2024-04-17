import { Main } from "@/components/main";
import { PageLabel } from "@/components/page-label";
import { ChainList } from "@/containers/chain/chain-list";

export default function Chains() {
  return (
    <Main>
      <PageLabel>Chains</PageLabel>
      <ChainList />
    </Main>
  );
}
