"use client";

import { ChainSelect } from "@/components/chain-select";
import { Pagination } from "@/components/pagination";
import { Table } from "@/components/table";
import { CHAIN_ID_NEUTRALITY } from "@/constants/chain";
import { useListBlocks } from "@/hooks/use-list-blocks";
import useChainState from "@/stores/use-chain-state";
import { formatAddress, formatTimestampFromNow } from "@/utils/formatter";
import Link from "next/link";
import { useState } from "react";
import Web3 from "web3";

export const BlockList = ({}) => {
  const chainID = useChainState((s) => s.id);
  const [page, setPage] = useState({ offset: 0, size: 10 });

  const listBlocks = useListBlocks({
    chainID,
    pageOffset: page.offset,
    pageSize: page.size,
  });

  return (
    <div className="mt-5">
      <Table
        className="max-md:min-w-[1000px]"
        label={`Total of ${BigInt(
          listBlocks.data?.totalSize ?? 0
        ).toLocaleString()} blocks`}
        headerComponent={
          <>
            <ChainSelect />
            <Pagination
              onOffsetChange={(offset) =>
                setPage((prev) => ({ ...prev, offset }))
              }
              offset={page.offset}
              size={page.size}
              totalSize={listBlocks.data?.totalSize ?? 0}
            />
          </>
        }
        data={listBlocks.data?.blocks ?? []}
        columns={[
          {
            label: "Block",
            render: (d) => (
              <Link className="text-[#0091C2]" href={"/blocks/" + d.height}>
                {d.height}
              </Link>
            ),
          },
          { label: "Age", render: (d) => formatTimestampFromNow(d.timestamp) },
          {
            label: "Txn",
            render: (d) => (
              <Link className="text-[#0091C2]" href={"/txs?block=" + d.height}>
                {d.txns}
              </Link>
            ),
          },
          {
            label: chainID === CHAIN_ID_NEUTRALITY ? "Validator" : "Signer",
            render: (d) => (
              <Link className="text-[#0091C2]" href={"/accounts/" + d.miner}>
                {formatAddress(d.miner)}
              </Link>
            ),
          },
          {
            label: "Gas Used",
            render: (d) => BigInt(d.gasUsed).toLocaleString(),
          },
          {
            label: "Gas Limit",
            render: (d) => BigInt(d.gasLimit).toLocaleString(),
          },
          // { label: "Base Fee", render: (d) => d.baseFeePerGas + " Gwei" },
          {
            label: "Reward",
            render: (d) => {
              return (
                Web3.utils
                  .fromWei(
                    d.txs?.reduce(
                      (p: any, c: any) =>
                        p + BigInt(c.effectiveGasPrice) * BigInt(c.gasUsed),
                      BigInt(0)
                    ),
                    "ether"
                  )
                  .toString() + " GEC"
              );
            },
          },
        ]}
        footerRightComponent={
          <Pagination
            onOffsetChange={(offset) =>
              setPage((prev) => ({ ...prev, offset }))
            }
            offset={page.offset}
            size={page.size}
            totalSize={listBlocks.data?.totalSize ?? 0}
          />
        }
      />
    </div>
  );
};
