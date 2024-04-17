"use client";

import { CarbonCard } from "@/components/carbon-card";
import {
  CHAIN_ID_EMISSION,
  CHAIN_ID_NEUTRALITY,
  CHAIN_ID_OFFSET,
  CHAIN_LABEL_EMISSION,
  CHAIN_LABEL_NEUTRALITY,
  CHAIN_LABEL_OFFSET,
} from "@/constants/chain";
import useChainState from "@/stores/use-chain-state";

export const HomeCarbonStats = ({}) => {
  const getChain = useChainState((s) => s.getChain);

  return (
    <div className="grid grid-cols-3 gap-x-5 mt-10 max-md:grid-cols-1 max-md:gap-x-0 max-md:gap-y-3">
      <CarbonCard
        srcBg="/neutral-bg.jpg"
        src="/neutral.png"
        tco2={
          getChain(CHAIN_ID_NEUTRALITY)?.carbonTotalAmount ?? "12000" ?? "0"
        }
        label={CHAIN_LABEL_NEUTRALITY}
      />
      <CarbonCard
        srcBg="/emission-bg.jpg"
        src="/emission.png"
        tco2={getChain(CHAIN_ID_EMISSION)?.carbonTotalAmount ?? "43000" ?? "0"}
        label={CHAIN_LABEL_EMISSION}
      />
      <CarbonCard
        srcBg="/offset-bg.jpg"
        src="/offset.png"
        tco2={getChain(CHAIN_ID_OFFSET)?.carbonTotalAmount ?? "55000" ?? "0"}
        label={CHAIN_LABEL_OFFSET}
      />
    </div>
  );
};
