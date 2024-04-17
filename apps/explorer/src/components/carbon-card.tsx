import Image from "next/image";

export const CarbonCard = ({
  label,
  tco2,
  src,
  srcBg,
}: {
  srcBg: string;
  src: string;
  label: string;
  tco2: string;
}) => {
  return (
    <div className="relative h-[146px] shadow-md rounded-[8px]">
      <Image
        src={srcBg}
        alt={label + " Bg"}
        layout="fill"
        objectFit="cover"
        objectPosition="center"
        className="rounded-[8px]"
      />
      <div className="flex flex-col p-6 absolute w-full z-10">
        <Image
          src={src}
          alt={label}
          width={25}
          height={30}
          objectFit="contain"
        />
        <span className="mt-4 text-[16px/24px] font-medium">{label}</span>
        <span className="text-[22px/28px] font-medium">
          {BigInt(tco2).toLocaleString()}
          <span className="text-[16px/20px] font-normal ml-1">tCO2</span>
        </span>
      </div>
    </div>
  );
};
