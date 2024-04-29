import { Main } from "@/components/main";
import { PageLabel } from "@/components/page-label";
import { TokenList } from "@/containers/token/list";

export default function NFTs() {
  return (
    <Main>
      <PageLabel>ERC721/ERC1155 Token List</PageLabel>
      <TokenList type="nft" />
    </Main>
  );
}
