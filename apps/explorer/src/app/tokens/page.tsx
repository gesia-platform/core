import { Main } from "@/components/main";
import { PageLabel } from "@/components/page-label";
import { TokenList } from "@/containers/token/list";

export default function Tokens() {
  return (
    <Main>
      <PageLabel>ERC20 Token List</PageLabel>
      <TokenList type="ft" />
    </Main>
  );
}
