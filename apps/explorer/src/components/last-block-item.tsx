import Image from "next/image";
import Link from "next/link";

export const LastBlockItem = ({
  time,
  number,
  proposer,
  txns,
}: {
  txns: number;
  proposer: string;
  number: number;
  time: string;
}) => {
  return (
    <div className="flex items-center py-2.5">
      <div className="rounded-[8px] bg-[#f5f6f7] w-[52px] h-[52px] flex">
        <Image
          src="/block.png"
          width={22}
          height={22}
          alt="Block"
          className="m-auto"
        />
      </div>

      <div className="ml-5 flex flex-col">
        <Link className="text-[14px] text-[#0091C2]" href={"/blocks/" + number}>
          {number}
        </Link>
        <span className="text-[12px] text-[#909396]">{time}</span>
      </div>

      <div className="ml-10 flex flex-col">
        <span className="text-[12px] text-[#1C1E20]">Validated By</span>
        <Link className="text-[14px] text-[#0091C2]" href={"/blocks/" + number}>
          {proposer}
        </Link>
      </div>
      <span className="ml-auto text-[14px] text-[#1c1e20]">{txns} txns</span>
    </div>
  );
};
