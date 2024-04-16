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
    <div className="py-[30px] px-5 flex flex-col h-full">
      <span className="text-[16px] font-medium mb-2.5">{label}</span>
      {children}
      <div className="mt-auto">
        <Link href={moreHref} className="mt-5 w-full">
          <button className="border rounded-[8px] border-[#eaeced] text-center text-[14px] h-[52px] w-full font-medium">
            {moreLabel}
          </button>
        </Link>
      </div>
    </div>
  );
};
