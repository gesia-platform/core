import { CarbonLabel } from "@/components/carbon-label";
import { Details } from "@/components/details";
import { DetailsRow } from "@/components/details-row";
import { DetailsRows } from "@/components/details-rows";
import { TxInputData } from "@/components/tx-input-data";
import { CHAIN_ID_NEUTRALITY, CHAIN_LABEL_NEUTRALITY } from "@/constants/chain";
import Image from "next/image";

export const TxDetails = ({}) => {
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
          <DetailsRow label="Transaction Hash">3123213123</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Result">#1</DetailsRow>
          <DetailsRow label="Block">#1</DetailsRow>
          <DetailsRow label="Timestamp">#1</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="From">#1</DetailsRow>
          <DetailsRow label="To">#1</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Value">#1</DetailsRow>
          <DetailsRow label="Total Tx Fee">#1</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Gas Price">#1</DetailsRow>
          <DetailsRow label="Gas Limit">#1</DetailsRow>
          <DetailsRow label="Gas Fees">#1</DetailsRow>
        </DetailsRows>

        <DetailsRows>
          <DetailsRow label="Input Data">
            <TxInputData />
          </DetailsRow>
        </DetailsRows>
      </Details>
    </div>
  );
};
