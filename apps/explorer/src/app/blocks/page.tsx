import { Main } from "@/components/main";
import { PageLabel } from "@/components/page-label";
import { BlockList } from "@/containers/block/list";

export default function Blocks() {
  return (
    <Main>
      <PageLabel>Blocks</PageLabel>
      <BlockList />
    </Main>
  );
}
