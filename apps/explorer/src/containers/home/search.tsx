import { SearchBar } from "@/components/search-bar";

export const HomeSearch = ({}) => {
  return (
    <div className="mt-2">
      <SearchBar
        inputProps={{
          placeholder: "Carbon Tx, Carbon Calculator Address, Wallet Address",
        }}
      />
    </div>
  );
};
