import { Main } from '@/components/main';
import { PageLabel } from '@/components/page-label';
import { CalculatorDetails } from '@/containers/calculator/details';

export default function CalculatorDetail({ params }: { params: any }) {
	if (params.chainID === 1) return;
	return (
		<Main>
			<PageLabel>Calculator Details</PageLabel>
			<CalculatorDetails voucherID={params.voucher} tokenNo={params.tokenNo} />
		</Main>
	);
}
