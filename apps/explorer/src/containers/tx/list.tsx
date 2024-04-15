import { ChainSelect } from "@/components/chain-select";
import { Pagination } from "@/components/pagination";
import { Table } from "@/components/table";

export const TxList = ({}) => {
  return (
    <div className="mt-5">
      <Table
        label={`More than 2,333,477,523 transactions found`}
        headerComponent={
          <>
            <ChainSelect />
            <Pagination totalPage={323} page={1} />
          </>
        }
        data={[{}, {}, {}, {}, {}, {}, {}, {}]}
        columns={[
          { label: "Txn Hash", render: () => "hih" },
          { label: "Block", render: () => "hih" },
          { label: "Age", render: () => "hih" },
          { label: "From", render: () => "hih" },
          { label: "To", render: () => "hih" },
          { label: "Value", render: () => "hih" },
          { label: "Txn Fee", render: () => "hih" },
        ]}
        footerRightComponent={<Pagination totalPage={323} page={1} />}
      />
    </div>
  );
};
