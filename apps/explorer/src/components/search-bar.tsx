import Image from "next/image";
import { DetailedHTMLProps, InputHTMLAttributes } from "react";

export const SearchBar = ({
  inputProps,
  buttonProps,
}: {
  inputProps?: DetailedHTMLProps<
    InputHTMLAttributes<HTMLInputElement>,
    HTMLInputElement
  >;
  buttonProps?: DetailedHTMLProps<
    React.ButtonHTMLAttributes<HTMLButtonElement>,
    HTMLButtonElement
  >;
}) => {
  return (
    <div className="rounded-[8px] bg-white p-[6px] flex items-center w-full max-w-[672px] border border-[#EAECED]">
      <input
        {...inputProps}
        className="w-full self-stretch placeholder:text-[#AAAEB1] text-[14px/24px] focus:outline-none ml-2 mr-4"
      />

      <button
        {...buttonProps}
        className="bg-[#00C1B3] rounded-[8px] flex items-center justify-center w-9 h-9"
      >
        <Image src="/search.png" alt="Search" width={15} height={15} />
      </button>
    </div>
  );
};
