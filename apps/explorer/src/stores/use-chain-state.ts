import { CHAIN_ID_NEUTRALITY } from "@/constants/chain";
import { create } from "zustand";

interface ChainState {
  id: number;
  setID: (id: number) => void;
}

const useChainState = create<ChainState>()((set) => ({
  id: CHAIN_ID_NEUTRALITY,
  setID: (id) => set((state) => ({ id })),
}));

export default useChainState;
