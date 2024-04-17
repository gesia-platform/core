"use client";

import {
  CHAIN_ID_EMISSION,
  CHAIN_ID_NEUTRALITY,
  CHAIN_ID_OFFSET,
} from "@/constants/chain";
import useChainState from "@/stores/use-chain-state";
import Image from "next/image";

export const CarbonLabel = ({ chainID }: { chainID: number }) => {
  const getChain = useChainState((s) => s.getChain);
  const chain = getChain(chainID);

  return (
    <div
      className={`flex items-center px-5 py-[15px] rounded-t-[8px]
    ${
      chain?.id === CHAIN_ID_NEUTRALITY
        ? "bg-[#332822]"
        : chain?.id === CHAIN_ID_EMISSION
          ? "bg-[#C24100]"
          : "bg-[#214240]"
    }`}
    >
      {chain?.id === CHAIN_ID_NEUTRALITY && (
        <Image src="/neutral-white.png" alt="Neutral" width={25} height={30} />
      )}
      {chain?.id === CHAIN_ID_EMISSION && (
        <Image
          src="/emission-white.png"
          alt="Emission"
          width={17}
          height={30}
        />
      )}
      {chain?.id === CHAIN_ID_OFFSET && (
        <Image src="/offset-white.png" alt="Offset" width={17} height={30} />
      )}

      <span className="ml-[15px] text-[16px] font-medium text-[#fff]">
        {chain?.label}
      </span>
    </div>
  );
};
