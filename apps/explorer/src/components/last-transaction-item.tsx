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
      <div className="rounded-[8px] bg-[#f5f6f7] w-[52px] h-[52px] flex">
        <Image
          src="/tx.png"
          width={22}
          height={22}
          alt="Block"
          className="m-auto"
        />
      </div>

      <div className="ml-5 flex flex-col">
        <div className="flex items-center">
          <span className="text-[12px] text-[#1C1E20] w-[46px]">Tx#</span>
          <Link
            className="text-[14px] text-[#0091C2]"
            href={"/txs/" + hash}
          >
            {hash}
          </Link>
        </div>

        <span className="text-[12px] text-[#909396]">{time}</span>
      </div>

      <div className="ml-10 flex flex-col">
        <div className="flex items-center">
          <span className="text-[12px] text-[#1C1E20] w-[46px]">From</span>
          <Link
            className="text-[14px] text-[#0091C2]"
            href={"/addresses/" + from}
          >
            {from}
          </Link>
        </div>
        <div className="flex items-center">
          <span className="text-[12px] text-[#1C1E20] w-[46px]">To</span>
          <Link className="text-[14px] text-[#0091C2]" href={"/addresses/" + to}>
            {to}
          </Link>
        </div>
      </div>
    </div>
  );
};
