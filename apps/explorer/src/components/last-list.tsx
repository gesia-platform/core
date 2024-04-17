import Link from "next/link";
import { ReactNode } from "react";

export const LastList = ({
  label,
  moreHref,
  moreLabel,
  children,
}: {
  children: ReactNode;
  label: string;
  moreLabel: string;
  moreHref: string;
}) => {
  return (
    <div className="last-of-type:border-l last-of-type:pl-5 flex flex-col border-[##EAECED] max-lg:last-of-type:!border-l-0 max-lg:last-of-type:!pl-0 max-lg:border-b max-lg:last-of-type:!border-b-0 max-lg:!pb-4 max-lg:pt-6">
      <span className="text-[16px] font-medium mb-2.5">{label}</span>
      {children}
      <div className="mt-auto">
        <Link href={moreHref} className="mt-5 w-full">
          <button className="border rounded-[8px] border-[#eaeced] text-center text-[14px] h-[52px] w-full">
            {moreLabel}
          </button>
        </Link>
      </div>
    </div>
  );
};
