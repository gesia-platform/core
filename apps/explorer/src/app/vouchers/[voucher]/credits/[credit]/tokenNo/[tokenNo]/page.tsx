import { Main } from '@/components/main';
import { PageLabel } from '@/components/page-label';
import { CreditDetails } from '@/containers/credit/details';

export default function CreditDetail({ params }: { params: any }) {
	if (params.chainID === 1) return;
	return (
		<Main>
			<PageLabel>Credit Details</PageLabel>
			<CreditDetails voucherID={params.voucher} creditID={params.credit} tokenNo={params.tokenNo} />
		</Main>
	);
}
