"use client";

import { useListChains } from "@/hooks/use-list-chains";
import useChainState from "@/stores/use-chain-state";
import { ReactNode, useEffect } from "react";

export default function ChainProvider({ children }: { children: ReactNode }) {
  const setChains = useChainState((s) => s.setChains);
  const chains = useListChains({});

  useEffect(() => {
    if (chains.data?.chains) {
      setChains(chains.data.chains);
    }
  }, [chains]);

  return children;
}
