"use client";

import { CarbonLabel } from "@/components/carbon-label";
import { Details } from "@/components/details";
import { DetailsRow } from "@/components/details-row";
import { DetailsRows } from "@/components/details-rows";
import { TxInputData } from "@/components/tx-input-data";
import { useGetTx } from "@/hooks/use-get-tx";
import useChainState from "@/stores/use-chain-state";
import { formatTimestamp, formatTimestampFromNow } from "@/utils/formatter";
import Link from "next/link";
import { redirect } from "next/navigation";
import { useMemo } from "react";
import Web3 from "web3";

export const TxDetails = ({ txID }: { txID: string }) => {
  const chainID = useChainState((s) => s.id);
  const getTx = useGetTx({ txID });
  const getChain = useChainState((s) => s.getChain);

  const tx = useMemo(() => {
    return getTx.data?.tx || null;
  }, [getTx.data]);

  const chain = useMemo(() => {
    const chainID = tx?.block?.chainID;
    return chainID ? getChain(chainID) : null;
  }, [getChain, tx]);

  if (getTx.error) return redirect("/error");
  
  if (!tx || !chain) return null;

  return (
    <div className="mt-5">
      <Details grid headerComponent={<CarbonLabel chainID={chainID} />}>
        <DetailsRows>
          <DetailsRow label="Chain ID">{`#${chain.id} ${chain.label}`}</DetailsRow>
          <DetailsRow label="Transaction Hash">{tx.hash}</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Status">
            {tx.status === "0" ? "Failure" : "Success"}
          </DetailsRow>
          <DetailsRow label="Block">
            <Link
              className="text-[#0091C2]"
              href={"/blocks/" + tx.block.height}
            >
              {tx.block.height}
            </Link>
          </DetailsRow>

          <DetailsRow label="Timestamp">
            {formatTimestampFromNow(tx.block.timestamp)} (
            {formatTimestamp(tx.block.timestamp)})
          </DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="From">
            <Link className="text-[#0091C2]" href={"/accounts/" + tx.from}>
              {tx.from}
            </Link>
          </DetailsRow>
          <DetailsRow label="To">
            <Link className="text-[#0091C2]" href={"/accounts/" + tx.to}>
              {tx.to}
            </Link>
          </DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Value">
            {Web3.utils.fromWei(tx.value, "ether") + " GEC"}
          </DetailsRow>
          <DetailsRow label="Transaction Fee">
            {Web3.utils
              .fromWei(
                BigInt(tx.effectiveGasPrice) * BigInt(tx.gasUsed),
                "ether"
              )
              .toString() + " GEC"}
          </DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Gas Price">
            {Web3.utils.fromWei(BigInt(tx.gasPrice), "ether").toString() +
              " GEC"}
          </DetailsRow>
          <DetailsRow label="Gas Limit">
            {BigInt(tx.gas).toLocaleString()}
          </DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Input Data">
            <TxInputData />
          </DetailsRow>
        </DetailsRows>
      </Details>
    </div>
  );
};
