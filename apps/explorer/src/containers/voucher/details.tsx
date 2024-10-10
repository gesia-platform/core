'use client';

import { CarbonLabel } from '@/components/carbon-label';
import { Details } from '@/components/details';
import { DetailsRow } from '@/components/details-row';
import { DetailsRows } from '@/components/details-rows';
import { Pagination } from '@/components/pagination';
import { Table } from '@/components/table';
import { useGetCalculators } from '@/hooks/use-get-calculators';
import { useGetVoucher } from '@/hooks/use-get-voucher';
import { useInfoTokens } from '@/hooks/use-info-tokens';
import { useListTokens } from '@/hooks/use-list-tokens';
import useChainState from '@/stores/use-chain-state';
import Link from 'next/link';
import { redirect } from 'next/navigation';
import { useEffect, useMemo, useState } from 'react';

export const VoucherDetails = ({ voucherID }: { voucherID: string }) => {
	const chainID = useChainState((s) => s.id);
	const getChain = useChainState((s) => s.getChain);
	const chain = getChain();
	const getVoucher = useGetVoucher({ voucherID, chainID });
	const [page, setPage] = useState({ offset: 0, size: 25 });
	const [mergedTokens, setMergedTokens] = useState<any[]>([]);

	const voucher = useMemo(() => {
		return getVoucher.data?.voucher || null;
	}, [getVoucher.data]);

	const listTokens = useListTokens({
		chainID,
		type: 'nft',
		pageOffset: page.offset,
		pageSize: page.size,
	});
	const infoTokens = useInfoTokens({ type: 'nft', chainID });

	const calculator = useGetCalculators({
		chainID,
		emissionContract: voucherID,
		pageOffset: page.offset,
		pageSize: page.size,
	});

	const mergeTokens = (arrayA: any[], arrayB: any[]) => {
		const map = new Map();

		arrayA.forEach((token: any) => map.set(token.address, { ...token }));

		arrayB.forEach((token: any) => {
			const key = token.address || token.nft_smart_contract;
			if (map.has(key)) {
				map.set(key, { ...map.get(key), ...token });
			}
		});

		return Array.from(map.values());
	};

	useEffect(() => {
		if (listTokens.data?.tokens.length > 0) {
			const tokens = mergeTokens(listTokens.data.tokens, infoTokens.data.list || []);
			console.log(tokens);
			setMergedTokens(tokens);
		} else {
			setMergedTokens([]);
		}
	}, [listTokens.data, infoTokens.data]);

	if (getVoucher.error) return redirect('/error');

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
					<DetailsRow label='Owner Address'>
						{voucher.owner ? (
							<Link className='text-[#0091C2]' href={'/accounts/' + voucher.address}>
								{voucher.owner}
							</Link>
						) : (
							'Owner does not exist'
						)}
					</DetailsRow>
				</DetailsRows>
			</Details>

			<div className='mt-5'>
				{chainID === 3 ? (
					<Table
						label='Voucher Token Holders'
						className='max-md:min-w-[1000px]'
						data={mergedTokens}
						columns={[
							{
								label: 'Token No',
								render: (d) => '# ' + d.voucher_token_id,
							},
							{
								label: 'Credit name',
								render: (d) => (
									<Link className='text-[#0091C2] inline-block w-[300px] truncate text-center' href={`${voucherID}/credits/${d.address}/tokenNo/${d.nft_token_id}`}>
										{d.nft_name_en}
									</Link>
								),
							},
							{
								label: 'Amount',
								render: (d) => BigInt(d.totalSupply).toLocaleString() + ' coc',
							},
							{
								label: 'Holders',
								render: (d) => d.holders,
							},
						]}
						footerRightComponent={<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={mergedTokens.length} />}
					/>
				) : chainID === 2 ? (
					<Table
						label='Emission Calculator Data'
						className='max-md:min-w-[1000px]'
						data={calculator.data?.content || []}
						columns={[
							{
								label: 'Token ID',
								render: (d) => '# ' + d.token_id,
							},
							{
								label: 'Calculator Name',
								render: (d) => (
									<Link className='text-[#0091C2] inline-block w-[300px] truncate text-center' href={`${voucherID}/calculator/tokenNo/${d.token_id}`}>
										{d.calculator_name + ' Calculator'}
									</Link>
								),
							},
							{
								label: 'Calculator Balance',
								render: (d) => d.calculator_balance.toLocaleString() + ' co2',
							},
							{
								label: 'Calculator Holders',
								render: (d) => d.calculator_holder.toLocaleString(),
							},
						]}
						footerRightComponent={
							<Pagination
								onOffsetChange={(offset) => {
									setPage((prev) => ({ ...prev, offset }));
								}}
								offset={page.offset}
								size={page.size}
								totalSize={calculator.data?.content.length}
							/>
						}
					/>
				) : (
					<div>No data available for the selected chain</div>
				)}
			</div>
		</div>
	);
};
