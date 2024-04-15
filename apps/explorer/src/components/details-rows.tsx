import { ReactNode } from "react";

export const DetailsRows = ({ children }: { children: ReactNode }) => {
  return (
    <div className="grid grid-flow-row row-start-auto gap-y-[15px]">
      {children}
    </div>
  );
};
