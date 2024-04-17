import { ReactNode } from "react";

export const DetailsRows = ({ children }: { children: ReactNode }) => {
  return (
    <div className="grid grid-flow-row gap-y-4 border-b border-b-[#EAECED] pb-[30px]">
      {children}
    </div>
  );
};
