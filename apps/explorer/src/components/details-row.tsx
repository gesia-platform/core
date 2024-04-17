import { ReactNode } from "react";

export const DetailsRow = ({
  label,
  children,
}: {
  label: ReactNode;
  children: ReactNode;
}) => {
  return (
    <div className="flex items-center max-md:flex-col max-md:!items-start">
      <div className="w-[194px] text-[#777a7d] text-[16px] max-md:!w-auto max-md:!text-[14px]">
        {label}:
      </div>
      <div className="pl-5 max-md:!pl-0 max-md:text-[14px] flex-1 relative">{children}</div>
    </div>
  );
};
