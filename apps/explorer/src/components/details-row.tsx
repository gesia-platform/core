import { ReactNode } from "react";

export const DetailsRow = ({
  label,
  children,
}: {
  label: ReactNode;
  children: ReactNode;
}) => {
  return (
    //max-md
    <div className="grid grid-flow-col auto-cols-min max-md:!grid-flow-row max-md:!auto-rows-auto">
      <div className="min-w-[194px] text-[#777a7d] text-[16px] max-md:!w-auto max-md:!text-[14px]">
        {label}:
      </div>
      <div className="pl-5 max-md:!pl-0 max-md:text-[14px] inline-block break-words whitespace-nowrap">
        {children}
      </div>
    </div>
  );
};
