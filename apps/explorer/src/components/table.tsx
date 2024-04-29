import { ReactNode } from "react";

export const Table = ({
  label,
  headerComponent,
  footerLeftComponent,
  footerRightComponent,
  columns,
  data,
  className,
}: {
  className: string;
  columns: (
    | { label: string; render: (data: any, index: number) => ReactNode }
    | null
    | undefined
  )[];
  data: any[];
  label?: string;
  headerComponent?: ReactNode;
  footerLeftComponent?: ReactNode;
  footerRightComponent?: ReactNode;
}) => {
  return (
    <div className={`bg-white border border-[#EAECED] shadow-md rounded-[8px]`}>
      {(label || headerComponent) && (
        <div className="flex items-center justify-between p-[15px] border-b-[#EAECED] border-b max-md:flex-col max-md:items-start">
          <span className="text-[16px] font-medium">{label}</span>
          <div className="max-md:mt-2 max-md:!grid max-md:!grid-flow-row max-md:auto-rows-auto max-md:gap-y-2 flex items-center">
            {headerComponent}
          </div>
        </div>
      )}

      <div className="overflow-auto">
        <table
          cellSpacing={0}
          cellPadding={0}
          className={`w-full ${className}`}
        >
          <thead className="">
            <tr>
              {columns.map((col, index) =>
                col ? (
                  <th
                    key={index}
                    className="h-[56px] text-[16px] font-medium px-2"
                  >
                    {col.label}
                  </th>
                ) : null
              )}
            </tr>
          </thead>
          <tbody>
            {data.map((x, i) => {
              return (
                <tr key={i + x._id}>
                  {columns.map((z, y) => {
                    if (!z) return null;
                    return (
                      <td
                        className="h-[52px] text-[16px] border-t-[#EAECED] border-t text-center px-2"
                        key={y}
                      >
                        {z.render(x, i)}
                      </td>
                    );
                  })}
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>

      {(footerLeftComponent || footerRightComponent) && (
        <div className="flex items-center justify-between p-[15px] border-t-[#EAECED] border-t max-md:flex-col max-md:items-stretch">
          <div>{footerLeftComponent}</div>
          <div className="flex items-center">{footerRightComponent}</div>
        </div>
      )}
    </div>
  );
};
