"use client";

import { AddressTag } from "@/components/address-tag";
import { CarbonLabel } from "@/components/carbon-label";
import { Details } from "@/components/details";
import { DetailsRow } from "@/components/details-row";
import { DetailsRows } from "@/components/details-rows";
import { TxInputData } from "@/components/tx-input-data";
import { useGetTx } from "@/hooks/use-get-tx";
import useChainState from "@/stores/use-chain-state";
import {
  formatGEC,
  formatTimestamp,
  formatTimestampFromNow,
} from "@/utils/formatter";
import Link from "next/link";
import { redirect } from "next/navigation";
import { useMemo } from "react";

export const TxDetails = ({ txID }: { txID: string }) => {
  const chainID = useChainState((s) => s.id);
  const getTx = useGetTx({ txID, chainID });
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
            <AddressTag address={tx.from} />
          </DetailsRow>
          <DetailsRow label="To">
            <Link className="text-[#0091C2]" href={"/accounts/" + tx.to}>
              {tx.to}
            </Link>
            <AddressTag address={tx.to} />
          </DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Value">{formatGEC(tx.value)}</DetailsRow>
          <DetailsRow label="Transaction Fee">
            {formatGEC(BigInt(tx.effectiveGasPrice) * BigInt(tx.gasUsed))}
          </DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Gas Price">
            {formatGEC(BigInt(tx.gasPrice))}
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
