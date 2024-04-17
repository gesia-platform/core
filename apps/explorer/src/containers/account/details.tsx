"use client";

import { CarbonLabel } from "@/components/carbon-label";
import { Details } from "@/components/details";
import { DetailsRow } from "@/components/details-row";
import { DetailsRows } from "@/components/details-rows";
import { useGetWeb3Account } from "@/hooks/use-get-web3-account";
import useChainState from "@/stores/use-chain-state";
import { useMemo } from "react";
import Web3 from "web3";

export const AccountDetails = ({ accountID }: { accountID: string }) => {
  const { id: chainID, getChain } = useChainState();
  const getWeb3Account = useGetWeb3Account({ accountID, chainID });

  const account = useMemo(() => {
    return getWeb3Account.data?.account || null;
  }, [getWeb3Account]);

  if (!account) return null;

  return (
    <div className="mt-5">
      <Details grid headerComponent={<CarbonLabel chainID={chainID} />}>
        <DetailsRows>
          <DetailsRow label="Chain ID">
            {`#${chainID} ${getChain()?.label}`}
          </DetailsRow>
          <DetailsRow label="Address">{account.address}</DetailsRow>

          <DetailsRow label="Balance">
            {Web3.utils.fromWei(BigInt(account.balance), "ether") + " GEC"}
          </DetailsRow>
        </DetailsRows>
      </Details>
    </div>
  );
};
