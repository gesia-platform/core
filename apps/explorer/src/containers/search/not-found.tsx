"use client";

import Image from "next/image";
import { useRouter } from "next/navigation";

export const SearchNotFound = ({
  general,
  string,
  error,
}: {
  error?: boolean;
  general?: boolean;
  string?: string;
}) => {
  const router = useRouter();

  return (
    <div className="h-[700px] max-md:!h-[50vh] relative w-full flex items-center">
      <Image
        src="/not-found.jpg"
        alt={"Not Found"}
        layout="fill"
        objectFit="cover"
        objectPosition="center"
        className="max-md:hidden"
      />

      <Image
        src="/not-found-m.jpg"
        alt={"Not Found"}
        layout="fill"
        objectFit="cover"
        objectPosition="center"
        className="hidden max-md:block"
      />

      <div className="absolute left-0 right-0 flex z-10 max-md:top-9">
        <div className="mx-auto w-full max-w-[1440px] px-[155px] max-md:!px-9 flex flex-col max-md:items-center">
          <span className="text-[42px] font-medium max-md:text-[32px]">
            {error
              ? "Sorry! We encountered an unexpected error."
              : general
                ? "Not found"
                : "Search not found"}
          </span>
          {!general && !error && (
            <span className="text-[16px] mt-5 max-md:text-center">
              Oops! The search string you entered was:{" "}
              <span className="font-medium">
                <br className="hidden max-md:block" />
                {string}
              </span>
              <br />
              Sorry! This is an invalid search string.
            </span>
          )}
          {error && (
            <span className="text-[16px] mt-5 max-md:text-center">
              An unexpected error occurred.
              <br />
              Please check back later
            </span>
          )}
          <button
            className="py-[15px] px-[50px] text-white rounded-[8px] text-center text-[14px] bg-[#00C1B3] self-start mt-10 max-md:!self-center"
            onClick={() => router.push("/")}
          >
            Back Home
          </button>
        </div>
      </div>
    </div>
  );
};
