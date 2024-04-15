import { Main } from "@/components/main";
import { BlockList } from "@/containers/blocks/block-list";

export default function Blocks() {
  return (
    <Main>
      <h2 className="text-[22px] font-medium text-[#1C1E20]">Blocks</h2>
      <BlockList />
    </Main>
  );
}
