"use client";

import { CarbonLabel } from "@/components/carbon-label";
import { Details } from "@/components/details";
import { DetailsRow } from "@/components/details-row";
import { DetailsRows } from "@/components/details-rows";
import { CHAIN_LABEL_NEUTRALITY } from "@/constants/chain";
import { useGetBlock } from "@/hooks/use-get-block";
import useChainState from "@/stores/use-chain-state";
import { formatTimestamp, formatTimestampFromNow } from "@/utils/formatter";
import Image from "next/image";
import Link from "next/link";
import { useMemo } from "react";

export const BlockDetails = ({ blockID }: { blockID: string }) => {
  const { id: chainID, getChain } = useChainState();
  const getBlock = useGetBlock({ blockID, chainID });

  const block = useMemo(() => {
    return getBlock.data?.block || null;
  }, [getBlock]);

  if (!block) return null;

  console.log(block);

  return (
    <div className="mt-5">
      <Details
        grid
        headerComponent={
          <CarbonLabel
            className="bg-[#332822] rounded-t-[8px]"
            label={`${CHAIN_LABEL_NEUTRALITY}`}
            icon={
              <Image
                src="/neutral-white.png"
                alt="Neutral"
                width={25}
                height={30}
              />
            }
          />
        }
      >
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
          {/**
           * <DetailsRow label="Proposed On">
            <Link className="text-[#0091C2]" href={"/addresses/" + block.miner}>
              {block.miner}
            </Link>
          </DetailsRow>
           */}
          <DetailsRow label="Transactions">
            {block.txns} transactions in this block
            {
              //and 8 contract internal transactions
            }
          </DetailsRow>
          {
            // <DetailsRow label="Withdrawals">{bloc}</DetailsRow>
          }
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Fee Recipient">
            <Link className="text-[#0091C2]" href={"/addresses/" + block.miner}>
              {block.miner}
            </Link>
          </DetailsRow>
          <DetailsRow label="Block Reward">#1</DetailsRow>
          <DetailsRow label="Total Difficulty">#1</DetailsRow>
          <DetailsRow label="Size">#1</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Gas Used">#1</DetailsRow>
          <DetailsRow label="Gas Limit">#1</DetailsRow>
          <DetailsRow label="Base Fee Per Gas">#1</DetailsRow>
          <DetailsRow label="Burnt Fees">#1</DetailsRow>
          <DetailsRow label="Extra Data">#1</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Hash">#1</DetailsRow>
          <DetailsRow label="Parent Hash">#1</DetailsRow>
          <DetailsRow label="StateRoot">#1</DetailsRow>
          <DetailsRow label="WithdrawalsRoot">#1</DetailsRow>
          <DetailsRow label="Nonce">#1</DetailsRow>
        </DetailsRows>
      </Details>
    </div>
  );
};
