import { Main } from '@/components/main';
import { PageLabel } from '@/components/page-label';
import { VoucherDetails } from '@/containers/voucher/details';

export default function VoucherDetail({ params }: { params: any }) {
	return (
		<Main>
			<PageLabel>Voucher Details</PageLabel>
			<VoucherDetails voucherID={params.voucher} />
		</Main>
	);
}
