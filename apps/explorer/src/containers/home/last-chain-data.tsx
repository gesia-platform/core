"use client";

import { LastBlockItem } from "@/components/last-block-item";
import { LastList } from "@/components/last-list";
import { LastTransactionItem } from "@/components/last-transaction-item";
import { CHAINS } from "@/constants/chain";
import { useState } from "react";

export const LastChainData = ({}) => {
  const [chain, setChain] = useState<string>("neutral");

  return (
    <div className="bg-white border border-[#eaeced] shadow-md rounded-[8px] mt-5">
      <div className="grid grid-cols-3">
        {CHAINS.map((data) => (
          <button
            onClick={() => {
              setChain(data.name);
            }}
            className={`h-[72px]  rounded-tl-[8px] font-medium text-[16px] ${
              chain === data.name
                ? "text-[#00C1B3] bg-white border-r"
                : "bg-[#f5f6f7]"
            }`}
          >
            Chain ID #{data.id} ({data.label})
          </button>
        ))}
      </div>

      <div className="grid grid-cols-2 gap-x-5">
        <LastList
          label="Last Blocks"
          moreHref="/blocks"
          moreLabel="VIEW ALL BLOCKS"
        >
          <LastBlockItem
            number={132132}
            txns={3213}
            time="11 secs ago"
            proposer="0x000000.00000000"
          />
          <LastBlockItem
            number={132132}
            txns={3213}
            time="11 secs ago"
            proposer="0x000000.00000000"
          />
          <LastBlockItem
            number={132132}
            txns={3213}
            time="11 secs ago"
            proposer="0x000000.00000000"
          />
          <LastBlockItem
            number={132132}
            txns={3213}
            time="11 secs ago"
            proposer="0x000000.00000000"
          />
          <LastBlockItem
            number={132132}
            txns={3213}
            time="11 secs ago"
            proposer="0x000000.00000000"
          />
        </LastList>

        <LastList
          label="Last Transactions"
          moreHref="/txs"
          moreLabel="VIEW ALL TRANSACTIONS"
        >
          <LastTransactionItem
            hash="0x00000...0000"
            time="10 sec ago"
            from="0x000000...0000"
            to="0x00000..000"
          />
          <LastTransactionItem
            hash="0x00000...0000"
            time="10 sec ago"
            from="0x000000...0000"
            to="0x00000..000"
          />
          <LastTransactionItem
            hash="0x00000...0000"
            time="10 sec ago"
            from="0x000000...0000"
            to="0x00000..000"
          />
          <LastTransactionItem
            hash="0x00000...0000"
            time="10 sec ago"
            from="0x000000...0000"
            to="0x00000..000"
          />
          <LastTransactionItem
            hash="0x00000...0000"
            time="10 sec ago"
            from="0x000000...0000"
            to="0x00000..000"
          />
        </LastList>
      </div>
    </div>
  );
};
