import { ReactNode } from "react";

export const Table = ({
  label,
  headerComponent,
  footerLeftComponent,
  footerRightComponent,
  columns,
  data,
}: {
  columns: { label: string; render: (data: any, index: number) => ReactNode }[];
  data: any[];
  label?: string;
  headerComponent?: ReactNode;
  footerLeftComponent?: ReactNode;
  footerRightComponent?: ReactNode;
}) => {
  return (
    <div className="bg-white border border-[#EAECED] shadow-md rounded-[8px]">
      {(label || headerComponent) && (
        <div className="flex items-center justify-between p-[15px] border-b-[#EAECED] border-b">
          <span className="text-[16px] font-medium">{label}</span>
          <div className="ml-auto flex items-center">{headerComponent}</div>
        </div>
      )}

      <table cellSpacing={0} cellPadding={0} className="w-full">
        <thead className="">
          <tr>
            {columns.map((col, index) => (
              <th
                key={index}
                className="h-[56px] text-[16px] font-medium "
              >
                {col.label}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data.map((x, i) => {
            return (
              <tr key={i}>
                {columns.map((z, y) => {
                  return (
                    <td
                      className="h-[52px] text-[16px] border-t-[#EAECED] border-t text-center"
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

      {(footerLeftComponent || footerRightComponent) && (
        <div className="flex items-center justify-between p-[15px] border-t-[#EAECED] border-t">
          <div>{footerLeftComponent}</div>
          <div className="ml-auto flex items-center">
            {footerRightComponent}
          </div>
        </div>
      )}
    </div>
  );
};
