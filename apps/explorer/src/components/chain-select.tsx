"use client";

import { CHAINS } from "@/constants/chain";
import useChainState from "@/stores/use-chain-state";
import Image from "next/image";

export const ChainSelect = ({}) => {
  const chainState = useChainState();

  return (
    <div className="mr-2 relative flex items-center border border-[#EAECED] rounded-[8px]">
      <select
        className="w-full text-[14px] h-[30px] px-[10px] cursor-pointer appearance-none pr-6"
        value={chainState.id.toString()}
        onChange={(e) => chainState.setID(Number(e.target.value))}
      >
        {CHAINS.map((x) => (
          <option key={x.name} value={x.id.toString()}>
            Chain ID #{x.id} ({x.label})
          </option>
        ))}
      </select>
      <div className="absolute right-2">
        <Image src="/bottom.png" width={10} height={5} alt="Bottom" />
      </div>
    </div>
  );
};
