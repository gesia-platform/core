import { Main } from "@/components/main";
import { PageLabel } from "@/components/page-label";
import { BlockDetails } from "@/containers/block/details";

export default function BlockDetail({ params }: { params: any }) {
  return (
    <Main>
      <PageLabel>Block Details</PageLabel>
      <BlockDetails blockID={params.block} />
    </Main>
  );
}
