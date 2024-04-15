import { CarbonLabel } from "@/components/carbon-label";
import { Details } from "@/components/details";
import { DetailsRow } from "@/components/details-row";
import { DetailsRows } from "@/components/details-rows";
import { CHAIN_ID_NEUTRALITY, CHAIN_LABEL_NEUTRALITY } from "@/constants/chain";
import Image from "next/image";

export const BlockDetails = ({}) => {
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
          <DetailsRow label="Chain ID">{`#${CHAIN_ID_NEUTRALITY} ${CHAIN_LABEL_NEUTRALITY}`}</DetailsRow>
          <DetailsRow label="Block Height">3123213123</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Timestamp">#1</DetailsRow>
          <DetailsRow label="Proposed On">#1</DetailsRow>
          <DetailsRow label="Transactions">#1</DetailsRow>
          <DetailsRow label="Withdrawals">#1</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Fee Recipient">#1</DetailsRow>
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
