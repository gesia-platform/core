import { SearchBar } from "@/components/search-bar";

export const HomeSearch = ({}) => {
  return (
    <div className="grid grid-flow-row gap-y-2 mt-5">
      <h1 className="text-[24px] leading-[32px] font-medium text-[#1C1E20]">
        GESIA Carbon Explorer
      </h1>

      <SearchBar
        inputProps={{
          placeholder: "Carbon Tx, Carbon Calculator Address, Wallet Address",
        }}
      />
    </div>
  );
};
