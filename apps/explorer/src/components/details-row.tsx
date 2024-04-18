import { ReactNode } from "react";

export const DetailsRow = ({
  label,
  children,
}: {
  label: ReactNode;
  children: ReactNode;
}) => {
  return (
    <div className="flex max-md:flex-col">
      <div className="min-w-[194px] text-[#777a7d] text-[16px] max-md:!w-auto max-md:!text-[14px]">
        {label}:
      </div>
      <div className="pl-5 max-md:!pl-0 max-md:text-[14px] break-all flex-1 max-md:flex-auto flex">
        {children}
      </div>
    </div>
  );
};
