import { formatAddress, formatHash } from "@/utils/formatter";
import Image from "next/image";
import Link from "next/link";

export const LastTransactionItem = ({
  time,
  hash,
  from,
  to,
}: {
  from: string;
  to: string;
  hash: string;
  time: string;
}) => {
  return (
    <div className="flex items-center py-2.5">
      <div className="rounded-[8px] bg-[#f5f6f7] min-w-[52px] min-h-[52px] flex">
        <div className="w-[22px] h-[22px] m-auto relative">
          <Image src="/tx.png" fill alt="Block" />
        </div>
      </div>

      <div className="ml-5 grid grid-flow-col auto-cols-auto gap-x-10 max-xl:gap-x-0 max-xl:!grid-flow-row  max-xl:flex-1">
        <div className="flex flex-col max-xl:!flex-row max-xl:justify-between">
          <div className="flex items-center">
            <span className="text-[12px] text-[#1C1E20] w-[46px]">Tx#</span>
            <Link className="text-[14px] text-[#0091C2]" href={"/txs/" + hash}>
              {formatHash(hash)}
            </Link>
          </div>

          <span className="text-[12px] text-[#909396]">{time}</span>
        </div>

        <div className="flex flex-col">
          <div className="flex items-center">
            <span className="text-[12px] text-[#1C1E20] w-[46px]">From</span>
            <Link
              className="text-[14px] text-[#0091C2]"
              href={"/accounts/" + from}
            >
              {formatAddress(from)}
            </Link>
          </div>
          <div className="flex items-center">
            <span className="text-[12px] text-[#1C1E20] w-[46px]">To</span>
            <Link
              className="text-[14px] text-[#0091C2]"
              href={"/accounts/" + to}
            >
              {formatAddress(to)}
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
};
