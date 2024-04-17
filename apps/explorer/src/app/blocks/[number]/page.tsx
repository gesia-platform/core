import { Main } from "@/components/main";
import { BlockDetails } from "@/containers/block/details";

export default function BlockDetail({ params }: { params: any }) {
  return (
    <Main>
      <h2 className="text-[22px] font-medium text-[#1C1E20]">Block Details</h2>
      <BlockDetails blockID={params.number} />
    </Main>
  );
}
