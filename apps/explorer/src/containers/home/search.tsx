'use client'

import { SearchBar } from "@/components/search-bar";
import { useRouter } from "next/navigation";

export const HomeSearch = ({}) => {
  const router = useRouter();
  return (
    <div className="mt-2">
      <SearchBar
        buttonProps={{
          onClick: () => {
            router.push("/search");
          },
        }}
        inputProps={{
          placeholder: "Carbon Tx, Carbon Calculator Address, Wallet Address",
          onKeyDown: (e) => {
            if (e.key === "Enter") {
              router.push("/search");
            }
          },
        }}
      />
    </div>
  );
};
