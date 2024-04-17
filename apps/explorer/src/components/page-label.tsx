import { ReactNode } from "react";

export const PageLabel = ({ children }: { children: ReactNode }) => {
  return <h2 className="text-[22px] max-md:!text-[20px] font-medium text-[#1C1E20]">{children}</h2>;
};
