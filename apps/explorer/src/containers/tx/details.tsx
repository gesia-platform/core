"use client";

import { AddressTag } from "@/components/address-tag";
import { CarbonLabel } from "@/components/carbon-label";
import { Details } from "@/components/details";
import { DetailsRow } from "@/components/details-row";
import { DetailsRows } from "@/components/details-rows";
import { ToAddress } from "@/components/to-address";
import { TxInputData } from "@/components/tx-input-data";
import { CHAIN_ID_EMISSION, CHAIN_VOUCHER_LOG_TOPIC_0_HASH } from "@/constants/chain";
import { useGetTx } from "@/hooks/use-get-tx";
import { useGetWeb3Account } from "@/hooks/use-get-web3-account";
import useChainState from "@/stores/use-chain-state";
import {
  formatGEC,
  formatTCO2,
  formatTimestamp,
  formatTimestampFromNow,
} from "@/utils/formatter";
import Link from "next/link";
import { redirect } from "next/navigation";
import { useMemo, useRef } from "react";
import Web3 from "web3";

export const TxDetails = ({ txID }: { txID: string }) => {
  const chainID = useChainState((s) => s.id);
  const getTx = useGetTx({ txID, chainID });
  const getChain = useChainState((s) => s.getChain);

  const tx = useMemo(() => {
    return getTx.data?.tx || null;
  }, [getTx.data]);

  const abi = useRef(new Web3().eth.abi).current;

  const chain = useMemo(() => {
    const chainID = tx?.block?.chainID;
    return chainID ? getChain(chainID) : null;
  }, [getChain, tx]);

  const getWeb3AccountTo = useGetWeb3Account(
    tx ? { accountID: tx.to, chainID } : undefined
  );

  const renderVoucher = () => {
    const log = tx.logs?.find(
      (x: any) => x.topics?.includes(CHAIN_VOUCHER_LOG_TOPIC_0_HASH)
    );

    if (!log) return null;

    const datas = log.data.replace("0x", "").match(/.{1,64}/g);
    const tokenID = BigInt(abi.decodeParameter("uint256", datas[0]) as any);
    const value = formatTCO2(abi.decodeParameter("uint256", datas[1]) as any,chainID === CHAIN_ID_EMISSION);

    return (
      <DetailsRow label="Carbon Voucher">
        <div>
          <div>
            <Link className="text-[#0091C2]" href={"/accounts/" + log.address}>
              {log.address}
            </Link>
            <span className="ml-2">(Contract)</span>
          </div>
          <div className="grid grid-flow-col auto-cols-max gap-x-2 items-center mt-4">
            <span>Token(App) ID: {tokenID.toString()}</span>
            <div className="h-4 w-[1px] bg-[#777A7D]" />
            <span>
              Amount: {(Number(value) / 1000).toString()}{" "}
              {chainID === 3 ? "tCOC" : "tCO2"}
            </span>
          </div>
        </div>
      </DetailsRow>
    );
  };
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
            <ToAddress tx={tx} tag />
          </DetailsRow>

          {getWeb3AccountTo.data?.account.isContract && renderVoucher()}
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
            <TxInputData input={tx.input} />
          </DetailsRow>
        </DetailsRows>
      </Details>
    </div>
  );
};
