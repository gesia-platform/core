import { ChainSelect } from "@/components/chain-select";
import { Pagination } from "@/components/pagination";
import { Table } from "@/components/table";

export const BlockList = ({}) => {
  return (
    <div className="mt-5">
      <Table
        label={`Total of 19,323,232 blocks`}
        headerComponent={
          <>
            <ChainSelect />
            <Pagination totalPage={323} page={1} />
          </>
        }
        data={[{}, {}, {}, {}, {}, {}, {}, {}]}
        columns={[
          { label: "Block", render: () => "hih" },
          { label: "Age", render: () => "hih" },
          { label: "Txn", render: () => "hih" },
          { label: "Fee Recipient", render: () => "hih" },
          { label: "Gas Used", render: () => "hih" },
          { label: "Gas Limit", render: () => "hih" },
          { label: "Base Fee", render: () => "hih" },
          { label: "Reward", render: () => "hih" },
          { label: "Brunt Fees (ETH)", render: () => "hih" },
        ]}
        footerRightComponent={<Pagination totalPage={323} page={1} />}
      />
    </div>
  );
};
