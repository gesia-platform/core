"use client";

import { useGetWeb3Account } from "@/hooks/use-get-web3-account";
import useChainState from "@/stores/use-chain-state";

export const AddressTag = ({
  address,
  creation,
}: {
  address: string;
  creation?: boolean;
}) => {
  const chainID = useChainState((s) => s.id);
  const getWeb3Account = useGetWeb3Account({ accountID: address, chainID });

  let result: string = "";
  if (getWeb3Account.data?.account) {
    if (getWeb3Account.data.account.isContract) {
      result = creation ? "Contract Creation" : "Contract";
    } else {
      if (getWeb3Account.data.account.isIOA) {
        result = "Internally Owned";
      } else {
        result = "Externally Owned";
      }
    }
  }

  return <span className="ml-2 block">({result}) </span>;
};
