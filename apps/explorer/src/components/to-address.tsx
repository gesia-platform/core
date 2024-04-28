import { formatAddress } from "@/utils/formatter";
import Link from "next/link";
import { useEffect, useState } from "react";
import { AddressTag } from "./address-tag";

export const ToAddress = ({
  tx,
  label,
  tag,
}: {
  tx: any;
  label?: boolean;
  tag?: boolean;
}) => {
  const [to, setTo] = useState("");

  useEffect(() => {
    if (tx.to) {
      setTo(tx.to);
    } else if (!tx.to && tx.contract) {
      setTo(tx.contract);
    }
  }, [tx]);

  if (!to) return "";

  const contractCreation = !tx.to;

  return (
    <div className="flex flex-row items-center justify-center">
      <Link className="text-[#0091C2]" href={"/accounts/" + to}>
        {label && contractCreation
          ? "Contract Creation"
          : tag
            ? to
            : formatAddress(to)}
      </Link>
      {tag && <AddressTag address={to} creation={contractCreation} />}
    </div>
  );
};
