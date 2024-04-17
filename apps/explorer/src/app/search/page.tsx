"use client";

import { SearchNotFound } from "@/containers/search/not-found";
import { redirect } from "next/navigation";

export default function Search({ searchParams }: { searchParams: any }) {
  const string = (searchParams.q || "") as string;

  if (string.length === 66) {
    // tx
    redirect("/txs/" + string);
  } else if (string.length === 42) {
    // account
    redirect("/accounts/" + string);
  } else if (!isNaN(Number(string))) {
    // block
    redirect("/blocks/" + string);
  }

  return <SearchNotFound string={string} />;
}
