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
        className={`py-[30px] max-md:!py-6  px-5 max-md:!px-4 ${
          grid
            ? "grid grid-flow-row gap-y-[30px] max-md:gap-x-6"
            : "flex flex-col"
        }`}
      >
        {children}
      </div>
    </div>
  );
};
