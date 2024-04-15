import { ChainSelect } from "@/components/chain-select";
import { Pagination } from "@/components/pagination";
import { Table } from "@/components/table";

export const ChainList = ({}) => {
  return (
    <div className="mt-5">
      <Table
        data={[{}, {}, {}]}
        columns={[
          { label: "#", render: () => "hih" },
          { label: "ID", render: () => "hih" },
          { label: "Chain", render: () => "hih" },
          { label: "Fee Consensus Algorithm", render: () => "hih" },
          { label: "Node Count", render: () => "hih" },
          { label: "Last Block Num", render: () => "hih" },
        ]}
      />
    </div>
  );
};
