export const TxInputData = ({ input }: { input: string }) => {
  const methodID = input?.substring(0, 10) ?? "";
  const inputs = input?.substring(10).match(/.{1,64}/g) ?? [];

  return (
    <div className="border border-[#EAECED] rounded-[8px] bg-[#FAFCFC] w-full">
      {methodID === "0x" ? (
        <div className="flex items-center px-5 py-2">{methodID}</div>
      ) : (
        <div className="flex items-center px-5 py-2">
          <span className="font-medium text-[16px] text-black">MethodID:</span>
          <span className="ml-5">{methodID}</span>
        </div>
      )}

      {inputs.length >= 1 && (
        <table
          cellPadding={0}
          cellSpacing={0}
          className="border-t border-t-[#EAECED] w-full"
        >
          <thead>
            <tr>
              <th className="px-5 py-2 text-[16px] border-b border-b-[#EAECED] font-medium text-center">
                #
              </th>
              <th className="px-5 py-2 text-[16px] border-b border-b-[#EAECED] font-medium text-left">
                Data
              </th>
            </tr>
          </thead>
          <tbody>
            {inputs.map((x, i) => (
              <tr key={i}>
                <td className="border-b border-b-[#EAECED] px-5 py-2 text-[16px] text-center">
                  {i}
                </td>
                <td className="border-b border-b-[#EAECED] px-5 py-2 text-[16px]">
                  {x}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </div>
  );
};
