import { Main } from "@/components/main";
import { TxList } from "@/containers/tx/list";
import Link from "next/link";

export default function Txs({ searchParams }: { searchParams: any }) {
  return (
    <Main>
      <h2 className="text-[22px] font-medium text-[#1C1E20]">Transactions</h2>
      {searchParams?.block && (
        <span className="text-[14px]">
          For Block{" "}
          <Link
            href={"/blocks/" + searchParams.block}
            className="text-[#0091C2] text-[16px]"
          >
            {searchParams.block}
          </Link>
        </span>
      )}
      <TxList blockID={searchParams?.block || undefined} />
    </Main>
  );
}
