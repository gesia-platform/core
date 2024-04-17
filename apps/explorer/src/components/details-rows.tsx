import { ReactNode } from "react";

export const DetailsRows = ({ children }: { children: ReactNode }) => {
  return (
    <div
      className={`grid grid-flow-row gap-y-4 max-md:!gap-y-3
    border-b last-of-type:!border-b-0 border-b-[#EAECED] 
    pb-[30px] last-of-type:!pb-0 max-md:!pb-6`}
    >
      {children}
    </div>
  );
};
