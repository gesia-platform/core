"use client";

import Image from "next/image";
import { useRouter } from "next/navigation";

export const SearchNotFound = ({}) => {
  const router = useRouter();

  return (
    <div className="h-[700px] relative w-full flex items-center">
      <Image
        src="/not-found.jpg"
        alt={"Not Found"}
        layout="fill"
        objectFit="cover"
        objectPosition="center"
        className="rounded-[8px]"
      />
      <div className="absolute left-0 right-0 flex z-10">
        <div className="mx-auto w-full max-w-[1440px] px-[155px] flex flex-col">
          <span className="text-[42px] font-medium">Search not found</span>
          <span className="text-[16px] mt-5">
            Oops! The search string you entered was:{" "}
            <span className="font-medium">Ddwd</span>
            <br />
            Sorry! This is an invalid search string.
          </span>
          <button
            className="py-[15px] px-[50px] text-white rounded-[8px] text-center text-[14px] bg-[#00C1B3] self-start mt-10"
            onClick={router.back}
          >
            Back Home
          </button>
        </div>
      </div>
    </div>
  );
};
