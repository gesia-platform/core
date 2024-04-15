import { ReactNode } from "react";

export const Details = ({
  children,
  headerComponent,
  grid,
}: {
  grid?: boolean;
  children: ReactNode;
  headerComponent?: ReactNode;
}) => {
  return (
    <div className="shadow-md border rounded-[8px] bg-white">
      {headerComponent}
      <div
        className={`py-[30px] px-5 ${
          grid
            ? "grid grid-flow-row auto-rows-auto gap-y-[25px]"
            : "flex flex-col"
        }`}
      >
        {children}
      </div>
    </div>
  );
};
