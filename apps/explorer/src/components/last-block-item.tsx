import { formatAddress } from "@/utils/formatter";
import Image from "next/image";
import Link from "next/link";

export const LastBlockItem = ({
  time,
  number,
  proposer,
  txns,
  minedLabel,
}: {
  minedLabel: string;
  txns: number;
  proposer: string;
  number: number;
  time: string;
}) => {
  return (
    <div className="flex items-center py-2.5">
      <div className="rounded-[8px] bg-[#f5f6f7] min-w-[52px] min-h-[52px] flex">
        <div className="w-[22px] h-[22px] relative m-auto">
          <Image src="/block.png" alt="Block" fill />
        </div>
      </div>

      <div className="ml-5 max-lg:!ml-3 max-lg:!min-w-[65px] flex flex-col min-w-[100px]">
        <Link className="text-[14px] text-[#0091C2]" href={"/blocks/" + number}>
          {number}
        </Link>
        <span className="text-[12px] text-[#909396]">{time}</span>
      </div>

      <div className="ml-10 max-xl:!ml-5 flex flex-col">
        <span className="text-[12px] text-[#1C1E20]">{minedLabel} By</span>
        <Link
          className="text-[14px] text-[#0091C2]"
          href={"/accounts/" + proposer}
        >
          {formatAddress(proposer)}
        </Link>
        <Link
          className="text-[14px] text-[#0091C2] hidden max-sm:!block"
          href={"/txs?block=" + number}
        >
          {txns} txns
        </Link>
      </div>
      <Link
        className="ml-auto text-[14px] text-[#0091C2] max-sm:hidden"
        href={"/txs?block=" + number}
      >
        {txns} txns
      </Link>
    </div>
  );
};
