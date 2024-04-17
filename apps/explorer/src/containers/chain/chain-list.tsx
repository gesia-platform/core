"use client";

import { Table } from "@/components/table";
import useChainState from "@/stores/use-chain-state";
import Link from "next/link";

export const ChainList = ({}) => {
  const setChainID = useChainState((s) => s.setID);
  const chains = useChainState((s) => s.chains);

  return (
    <div className="mt-5">
      <Table
        data={chains}
        columns={[
          { label: "Chain ID", render: (d) => d.id },
          { label: "Chain Name", render: (d) => d.label },
          {
            label: "Consensus Algorithm",
            render: (d) => d.consensusAlgorithm,
          },
          { label: "Nodes", render: (d) => d.nodes },
          {
            label: "Last Block",
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
