"use client";

import { Table } from "@/components/table";
import { useListChains } from "@/hooks/use-list-chains";
import useChainState from "@/stores/use-chain-state";
import Link from "next/link";

export const ChainList = ({}) => {
  const listChains = useListChains({});
  const setChainID = useChainState((s) => s.setID);

  return (
    <div className="mt-5">
      <Table
        data={listChains.data?.chains ?? []}
        columns={[
          { label: "Chain ID", render: (d) => d.id },
          { label: "Chain Name", render: (d) => d.label },
          {
            label: "Fee Consensus Algorithm",
            render: (d) => d.consensusAlgorithm,
          },
          { label: "Node Count", render: (d) => d.nodes },
          {
            label: "Last Block Num",
            render: (d) => (
              <Link
                className=" text-[#0091C2]"
                href={"/blocks/" + d.latestBlockHeight}
                onClick={() => {
                  setChainID(d.id);
                }}
              >
                {d.latestBlockHeight}
              </Link>
            ),
          },
        ]}
      />
    </div>
  );
};
