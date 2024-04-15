import Image from "next/image";

export const Pagination = ({
  totalPage,
  page,
}: {
  totalPage: number;
  page: number;
}) => {
  return (
    <div className="grid grid-flow-col auto-cols-auto gap-x-2">
      <button className="h-[30px] px-[10px] border border-[#EAECED] text-[14px] text-[#1C1E20] rounded-[8px]">
        First
      </button>

      <button className="w-[30px] h-[30px] flex border border-[#EAECED] rounded-[8px]">
        <Image className="m-auto" width={5} height={10} alt="Left" src={"/left.png"} />
      </button>

      <div className="h-[30px] flex border border-[#EAECED]  text-[14px] text-[#1C1E20] rounded-[8px] px-[10px] flex items-center">
        {page} of {totalPage}
      </div>

      <button className="w-[30px] h-[30px] flex border border-[#EAECED] rounded-[8px]">
        <Image className="m-auto" width={5} height={10} alt="Right" src={"/right.png"} />
      </button>

      <button className="h-[30px] px-[10px] border border-[#EAECED] text-[14px] text-[#1C1E20] rounded-[8px]">
        Last
      </button>
    </div>
  );
};
