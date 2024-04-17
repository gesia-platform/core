"use client";

import Image from "next/image";
import { useEffect, useState } from "react";

export const CarbonCard = ({
  label,
  tco2,
  src,
  srcBg,
  smallIcon,
}: {
  smallIcon?: boolean;
  srcBg: string;
  src: string;
  label: string;
  tco2: string;
}) => {
  const [isClient, setIsClient] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, [isClient]);

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
          width={smallIcon ? 17 : 25}
          height={30}
          objectFit="contain"
        />
        <span className="mt-4 text-[16px] font-medium">{label}</span>
        {isClient && (
          <span className="text-[22px] font-medium">
            {BigInt(tco2).toLocaleString()}
            <span className="text-[16px] font-normal ml-1">tCO2</span>
          </span>
        )}
      </div>
    </div>
  );
};
