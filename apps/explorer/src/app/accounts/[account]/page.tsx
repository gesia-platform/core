import { Main } from "@/components/main";
import { AccountDetails } from "@/containers/account/details";

export default function AccountsDetail({ params }: { params: any }) {
  return (
    <Main>
      <h2 className="text-[22px] font-medium text-[#1C1E20]">
        Account Details
      </h2>
      <AccountDetails accountID={params.account} />
    </Main>
  );
}
