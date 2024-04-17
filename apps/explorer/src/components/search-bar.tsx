import Image from "next/image";
import { DetailedHTMLProps, InputHTMLAttributes } from "react";
import { ChainSelect } from "./chain-select";

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
    <div className="rounded-[8px] bg-white p-[6px] flex items-center w-full max-w-[780px] border border-[#EAECED]">
      <ChainSelect borderHidden />
      <input
        {...inputProps}
        className="flex-1 self-stretch placeholder:text-[#AAAEB1] text-[14px] focus:outline-none ml-2 mr-4"
      />

      <button
        {...buttonProps}
        className="bg-[#00C1B3] rounded-[8px] flex items-center justify-center w-9 h-9"
      >
        <div className="w-[15px] h-[15px] m-auto relative">
          <Image src="/search.png" alt="Search" fill />
        </div>
      </button>
    </div>
  );
};
