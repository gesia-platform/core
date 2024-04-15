import { CHAINS } from "@/constants/chain";
import Image from "next/image";

export const ChainSelect = ({}) => {
  return (
    <div className="mr-2 relative flex items-center">
      <select className="border border-[#EAECED] rounded-[8px] text-[14px] h-[30px] px-[10px] cursor-pointer appearance-none pr-6">
        {CHAINS.map((x) => (
          <option key={x.name}>
            Chain ID #{x.id} ({x.label})
          </option>
        ))}
      </select>
      <div className="absolute right-2">
        <Image src="/bottom.png" width={10} height={5} alt="Bottom" />
      </div>
    </div>
  );
};
