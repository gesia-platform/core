import { CarbonCard } from "@/components/carbon-card";
import {
  CHAIN_LABEL_EMISSION,
  CHAIN_LABEL_NEUTRALITY,
  CHAIN_LABEL_OFFSET,
} from "@/constants/chain";

export const HomeCarbonStats = ({}) => {
  return (
    <div className="grid grid-cols-3 gap-x-5 mt-10">
      <CarbonCard
        srcBg="/neutral-bg.jpg"
        src="/neutral.png"
        tco2={-52000}
        label={CHAIN_LABEL_NEUTRALITY}
      />
      <CarbonCard
        srcBg="/emission-bg.jpg"
        src="/emission.png"
        tco2={12000}
        label={CHAIN_LABEL_EMISSION}
      />
      <CarbonCard
        srcBg="/offset-bg.jpg"
        src="/offset.png"
        tco2={3000}
        label={CHAIN_LABEL_OFFSET}
      />
    </div>
  );
};
