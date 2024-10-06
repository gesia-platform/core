'use client';

import { CarbonLabel } from '@/components/carbon-label';
import { Details } from '@/components/details';
import { DetailsRow } from '@/components/details-row';
import { DetailsRows } from '@/components/details-rows';
import { Pagination } from '@/components/pagination';
import { Table } from '@/components/table';
import { useGetCredit } from '@/hooks/use-get-credit';
import { useGetVoucher } from '@/hooks/use-get-voucher';
import { Spinner } from '@/components/spinner';
import useChainState from '@/stores/use-chain-state';
import { formatAddress, formatHash, formatTimestampFromNow } from '@/utils/formatter';
import Link from 'next/link';
import { redirect } from 'next/navigation';
import { useEffect, useMemo, useState } from 'react';
import { useInfoToken } from '@/hooks/use-info-token';

export const CreditDetails = ({ voucherID, creditID, tokenNo }: { voucherID: string; creditID: string; tokenNo: string }) => {
	const chainID = useChainState((s) => s.id);
	const [page, setPage] = useState({ offset: 0, size: 10 });
	const [isLoading, setIsLoading] = useState(true);

	const getChain = useChainState((s) => s.getChain);
	const chain = getChain();
	const getVoucher = useGetVoucher({ voucherID, chainID });
	const getCredit = useGetCredit({ creditID, chainID, tokenNo, pageOffset: page.offset, pageSize: page.size });
	const getCreditName = useInfoToken({ type: 'nft', chainID, nft_smart_contract: creditID });

	const voucher = useMemo(() => {
		return getVoucher.data?.voucher || null;
	}, [getVoucher.data]);

	const credits = useMemo(() => {
		return getCredit.data || [];
	}, [getCredit.data]);

	const paginatedCredits = useMemo(() => {
		return credits.slice(page.offset * page.size, (page.offset + 1) * page.size);
	}, [credits, page]);

	useEffect(() => {
		if (credits.length === 0) {
			setIsLoading(true);
		} else {
			setIsLoading(false);
		}
	}, [credits, page]);

	if (getCredit.error) return redirect('/error');

	if (!voucher || !chain) return null;

	return (
		<div className='mt-5'>
			<Details grid headerComponent={<CarbonLabel chainID={chainID} />}>
				<DetailsRows>
					<DetailsRow label={chain.label}>{`Chain ID - ${chain.id}`}</DetailsRow>
					<DetailsRow label='Registry name'>{voucher.name}</DetailsRow>
					<DetailsRow label='Token Type'>{voucher.type}</DetailsRow>
					<DetailsRow label='Voucher Address'>
						<Link className='text-[#0091C2]' href={'/accounts/' + voucher.address}>
							{voucher.address}
						</Link>
					</DetailsRow>
					<DetailsRow label='Owner Address'>Owner does not exist</DetailsRow>
					<DetailsRow label='Token ID'># {tokenNo}</DetailsRow>
					<DetailsRow label='Credit Name'>{getCreditName.data?.nft_name_en}</DetailsRow>
				</DetailsRows>
			</Details>
			<div className='mt-5'>
				{isLoading ? (
					<div className='flex justify-center items-center h-32'>
						<Spinner size={8} />
					</div>
				) : (
					<Table
						label='Transaction List'
						className='max-md:min-w-[1000px]'
						data={paginatedCredits}
						columns={[
							{
								label: 'Transaction Hash',
								render: (d) => (
									<Link className='text-[#0091C2]' href={'/txs/' + d.transactionHash}>
										{formatHash(d.transactionHash)}
									</Link>
								),
							},
							{
								label: 'Method',
								render: (d) => <div>{d.methodName}</div>,
							},
							{
								label: 'Block',
								render: (d) => <div>{d.blockNumber}</div>,
							},
							{
								label: 'Age',
								render: (d) => <div>{formatTimestampFromNow(d.creationTime)}</div>,
							},
							{
								label: 'From',
								render: (d) => (
									<Link className='text-[#0091C2]' href={'/accounts/' + d.from}>
										{formatAddress(d.from)}
									</Link>
								),
							},
							{
								label: 'To',
								render: (d) => (
									<Link className='text-[#0091C2]' href={'/accounts/' + d.to}>
										{formatAddress(d.to)}
									</Link>
								),
							},
							{
								label: 'Amount',
								render: (d) => <div>{BigInt(d.amount).toLocaleString() + ' coc'}</div>,
							},
						]}
						footerRightComponent={<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={credits.length} />}
					/>
				)}
			</div>
		</div>
	);
};
