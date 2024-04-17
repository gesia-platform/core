import { Main } from "@/components/main";
import { PageLabel } from "@/components/page-label";
import { AccountDetails } from "@/containers/account/details";

export default function AccountsDetail({ params }: { params: any }) {
  return (
    <Main>
      <PageLabel>Account Details</PageLabel>
      <AccountDetails accountID={params.account} />
    </Main>
  );
}
