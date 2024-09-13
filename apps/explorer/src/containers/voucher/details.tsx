'use client';

import { CarbonLabel } from '@/components/carbon-label';
import { Details } from '@/components/details';
import { DetailsRow } from '@/components/details-row';
import { DetailsRows } from '@/components/details-rows';
import { Pagination } from '@/components/pagination';
import { Table } from '@/components/table';
import { useGetVoucher } from '@/hooks/use-get-voucher';
import { useListVoucherTokens } from '@/hooks/use-list-voucher-tokens';
import useChainState from '@/stores/use-chain-state';
import Link from 'next/link';
import { redirect } from 'next/navigation';
import { useMemo, useState } from 'react';

export const VoucherDetails = ({ voucherID }: { voucherID: string }) => {
	const chainID = useChainState((s) => s.id);
	const getVoucher = useGetVoucher({ voucherID, chainID });
	const getChain = useChainState((s) => s.getChain);
	const [page, setPage] = useState({ offset: 0, size: 10 });

	const voucher = useMemo(() => {
		return getVoucher.data?.voucher || null;
	}, [getVoucher.data]);

	const chain = getChain();

	const listVoucherTokens = useListVoucherTokens(
		voucher
			? {
					chainID,
					voucherID,
				}
			: undefined,
	);

	const carbonCreditsTokenIds = useMemo(() => {
		const envTokenIds = process.env.NEXT_PUBLIC_CARBON_CREDITS_TOKEN_IDS || '';
		return envTokenIds.split(',').map((id) => id.trim());
	}, []);

	const filteredTokens = useMemo(() => {
		return (
			listVoucherTokens.data?.tokens?.filter((token: any) => {
				const tokenId = token.id.toString();
				return carbonCreditsTokenIds.includes(tokenId);
			}) ?? []
		);
	}, [listVoucherTokens.data, carbonCreditsTokenIds]);

	const tokens = useMemo(() => {
		return filteredTokens.slice(page.offset, page.offset + page.size);
	}, [page, filteredTokens]);

	if (getVoucher.error) return redirect('/error');

	if (!voucher || !chain) return null;

	return (
		<div className='mt-5'>
			<Details grid headerComponent={<CarbonLabel chainID={chainID} />}>
				<DetailsRows>
					<DetailsRow label='Chain ID'>{`#${chain.id} ${chain.label}`}</DetailsRow>
					<DetailsRow label='Voucher Name'>{voucher.name}</DetailsRow>
					<DetailsRow label='Voucher Address'>
						<Link className='text-[#0091C2]' href={'/accounts/' + voucher.address}>
							{voucher.address}
						</Link>
					</DetailsRow>
				</DetailsRows>
			</Details>

			<div className='mt-5'>
				<Table
					label='Voucher Token Holders'
					className='max-md:min-w-[1000px]'
					headerComponent={
						<>
							<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={filteredTokens.length} />
						</>
					}
					data={tokens}
					columns={[
						{
							label: 'ID',
							render: (d) => '#' + d.id,
						},
						{
							label: 'Oracle Owner',
							render: (d) => d.owner,
						},
						{
							label: 'Quantity',
							render: (d) => d.amount,
						},
					]}
					footerRightComponent={<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={filteredTokens.length} />}
				/>
			</div>
		</div>
	);
};
