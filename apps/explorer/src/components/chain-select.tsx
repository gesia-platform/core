"use client";

import {
  CHAINS,
  CHAIN_ID_EMISSION,
  CHAIN_ID_NEUTRALITY,
} from "@/constants/chain";
import useChainState from "@/stores/use-chain-state";
import Image from "next/image";
import { useEffect } from "react";

export const ChainSelect = ({
  borderHidden,
  worker,
}: {
  borderHidden?: boolean;
  worker?: boolean;
}) => {
  const chainState = useChainState();
  useEffect(() => {
    if (worker && chainState.id === CHAIN_ID_NEUTRALITY) {
      chainState.setID(CHAIN_ID_EMISSION);
    }
  }, [worker, chainState.id]);

  return (
    <div
      className={`mr-2 relative flex items-center  rounded-[8px] ${
        !borderHidden && "border border-[#EAECED]"
      }`}
    >
      <select
        className="w-full text-[14px] h-[30px] px-[10px] cursor-pointer appearance-none pr-6 bg-transparent  rounded-[8px]"
        value={chainState.id.toString()}
        onChange={(e) => chainState.setID(Number(e.target.value))}
      >
        {CHAINS.map((x) =>
          worker && x.id === CHAIN_ID_NEUTRALITY ? null : (
            <option key={x.name} value={x.id.toString()}>
              Chain ID #{x.id} ({x.label})
            </option>
          )
        )}
      </select>
      <div className="absolute right-2">
        <Image src="/bottom.png" width={10} height={5} alt="Bottom" />
      </div>
    </div>
  );
};
