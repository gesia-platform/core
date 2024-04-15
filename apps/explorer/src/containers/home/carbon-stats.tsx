import { CarbonCard } from "@/components/carbon-card";

export const HomeCarbonStats = ({}) => {
  return (
    <div className="grid grid-cols-3 gap-x-5 mt-10">
      <CarbonCard
        srcBg="/emission-bg.jpg"
        src="/emission.png"
        tco2={12000}
        label="Emission"
      />
      <CarbonCard
        srcBg="/offset-bg.jpg"
        src="/offset.png"
        tco2={3000}
        label="Offset"
      />
      <CarbonCard
        srcBg="/neutral-bg.jpg"
        src="/neutral.png"
        tco2={-52000}
        label="Neutral"
      />
    </div>
  );
};
