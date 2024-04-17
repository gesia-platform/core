import { ReactNode } from "react";

export const DetailsRow = ({
  label,
  children,
}: {
  label: ReactNode;
  children: ReactNode;
}) => {
  return (
    <div className="flex items-center">
      <div className="w-[194px] text-[#777a7d] text-[16px]">{label}:</div>
      <div className="pl-5 flex-1 relative">
          {children}
      </div>
    </div>
  );
};
