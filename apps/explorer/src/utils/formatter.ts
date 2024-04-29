import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";
import utc from "dayjs/plugin/utc";
import Web3 from "web3";
dayjs.extend(relativeTime);
dayjs.extend(utc);

export const formatAddress = (address: string) =>
  address?.substring(0, 13) + "..." + address?.substring(address.length - 10);

export const formatHash = (hash: string) => hash.substring(0, 15) + "...";

export const formatTCO2 = (value: string, isEmission: boolean) => {
  return isEmission
    ? Number(BigInt(value) / BigInt(1000000000)).toFixed(2)
    : Number(BigInt(value)).toFixed(2);
};
export const formatTimestampFromNow = (timestamp: string) => {
  let result = dayjs.unix(Number(timestamp)).local().fromNow(true) + " ago";

  result = result.replace(" few ", " ");

  return result;
};

export const formatTimestamp = (timestamp: string) => {
  return dayjs.unix(Number(timestamp)).local().toString();
};

export const formatGEC = (data: bigint | undefined | null) => {
  let result = Web3.utils.fromWei(data || BigInt(0), "ether");
  if (result === "0.") {
    result = result.replace(".", "");
  }

  return result + " GEC";
};
