"use client";

import { CarbonLabel } from "@/components/carbon-label";
import { Details } from "@/components/details";
import { DetailsRow } from "@/components/details-row";
import { DetailsRows } from "@/components/details-rows";
import { useGetBlock } from "@/hooks/use-get-block";
import useChainState from "@/stores/use-chain-state";
import {
  formatNZC,
  formatTimestamp,
  formatTimestampFromNow,
} from "@/utils/formatter";
import Link from "next/link";
import { redirect } from "next/navigation";
import { useMemo } from "react";
import Web3 from "web3";

export const BlockDetails = ({ blockID }: { blockID: string }) => {
  const { id: chainID, getChain } = useChainState();
  const getBlock = useGetBlock({ blockID, chainID });

  const block = useMemo(() => {
    return getBlock.data?.block || null;
  }, [getBlock]);

  if (getBlock.error) return redirect("/error");

  if (!block) return null;

  return (
    <div className="mt-5">
      <Details grid headerComponent={<CarbonLabel chainID={chainID} />}>
        <DetailsRows>
          <DetailsRow label="Chain ID">{`#${chainID} ${getChain()
            ?.label}`}</DetailsRow>
          <DetailsRow label="Block Height">{block.height}</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Timestamp">
            {formatTimestampFromNow(block.timestamp)} (
            {formatTimestamp(block.timestamp)})
          </DetailsRow>
          <DetailsRow label="Transactions">
            <Link
              className="text-[#0091C2]"
              href={"/txs?block=" + block.height}
            >
              {block.txns} transactions
            </Link>
            {" "}in this block
          </DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow
            label={
              getChain()?.consensusAlgorithm === "PoS" ? "Validator" : "Signer"
            }
          >
            <Link className="text-[#0091C2]" href={"/accounts/" + block.miner}>
              {block.miner}
            </Link>
          </DetailsRow>
          <DetailsRow label="Block Reward">
            {formatNZC(
              block.txs?.reduce(
                (p: any, c: any) =>
                  p + BigInt(c.effectiveGasPrice) * BigInt(c.gasUsed),
                BigInt(0)
              )
            )}
          </DetailsRow>
          <DetailsRow label="Total Difficulty">
            {BigInt(block.totalDifficulty).toLocaleString()}
          </DetailsRow>
          <DetailsRow label="Size">
            {BigInt(block.size).toLocaleString()} bytes
          </DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Gas Used">
            {BigInt(block.gasUsed).toLocaleString()}
          </DetailsRow>
          <DetailsRow label="Gas Limit">
            {BigInt(block.gasLimit).toLocaleString()}
          </DetailsRow>
          {Boolean(block.baseFeePerGas) && (
            <DetailsRow label="Base Fee Per Gas">
              {BigInt(block.baseFeePerGas).toLocaleString()}
            </DetailsRow>
          )}
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Hash">{block.hash}</DetailsRow>
          <DetailsRow label="Parent Hash">{block.parentHash}</DetailsRow>
          <DetailsRow label="StateRoot">{block.stateRoot}</DetailsRow>
          <DetailsRow label="ReceiptsRoot">{block.receiptsRoot}</DetailsRow>
          <DetailsRow label="TransactionsRoot">
            {block.transactionsRoot}
          </DetailsRow>
          {Boolean(block.withdrawalsRoot) && (
            <DetailsRow label="WithdrawalsRoot">
              {block.withdrawalsRoot}
            </DetailsRow>
          )}
          <DetailsRow label="Nonce">
            {block.nonce === "0" ? "0x0000000000000000" : block.nonce}
          </DetailsRow>
        </DetailsRows>
      </Details>
    </div>
  );
};
