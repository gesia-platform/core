"use client";

import { SearchBar } from "@/components/search-bar";
import { useRouter } from "next/navigation";
import { useRef } from "react";

export const HomeSearch = ({}) => {
  const router = useRouter();

  const input = useRef<HTMLInputElement>(null);

  return (
    <div className="mt-2">
      <SearchBar
        buttonProps={{
          onClick: () => {
            router.push(
              "/search?string=" + encodeURIComponent(input.current?.value ?? "")
            );
          },
        }}
        inputProps={{
          ref: input,
          onChange: (e) => {
            if (input.current) {
              input.current.value = e.target.value;
            }
          },
          placeholder:
            "Search by Address / Block / Txn Hash / Voucher / Calculator",
          onKeyDown: (e) => {
            if (e.key === "Enter") {
              router.push(
                "/search?string=" +
                  encodeURIComponent(input.current?.value ?? "")
              );
            }
          },
        }}
      />
    </div>
  );
};
