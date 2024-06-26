import { CHAIN_ID_NEUTRALITY } from "@/constants/chain";
import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

interface Chain {
  id: number;
  name: string;
  label: string;
  consensusAlgorithm: string;
  latestBlockHeight: string;
  carbonTotalAmount: string;
  nodes: number;
}

interface ChainState {
  id: number;
  setID: (id: number) => void;
  chains: Chain[];
  setChains: (chains: Chain[]) => void;
  getChain: (id?: number) => Chain | undefined;
}

const useChainState = create<ChainState>()(
  persist(
    (set, get) => ({
      id: CHAIN_ID_NEUTRALITY,
      setID: (id) => set((state) => ({ id })),
      chains: [],
      setChains: (chains) => set({ chains }),
      getChain: (id) => {
        return get().chains.filter((x) => x.id === (id || get().id))[0];
      },
    }),
    {
      name: "chain-storage",
      storage: createJSONStorage(() => localStorage),
    }
  )
);

export default useChainState;
