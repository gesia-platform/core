'use client';

import { CarbonLabel } from '@/components/carbon-label';
import { Details } from '@/components/details';
import { DetailsRow } from '@/components/details-row';
import { DetailsRows } from '@/components/details-rows';
import { Pagination } from '@/components/pagination';
import { Table } from '@/components/table';
import { CHAIN_ID_OFFSET } from '@/constants/chain';
import { useGetVoucher } from '@/hooks/use-get-voucher';
import { useInfoTokens } from '@/hooks/use-info-tokens';
import { useListTokens } from '@/hooks/use-list-tokens';
import { useListVoucherTokens } from '@/hooks/use-list-voucher-tokens';
import useChainState from '@/stores/use-chain-state';
import Link from 'next/link';
import { redirect } from 'next/navigation';
import { useEffect, useMemo, useState } from 'react';

export const VoucherDetails = ({ voucherID }: { voucherID: string }) => {
	const chainID = useChainState((s) => s.id);
	const getVoucher = useGetVoucher({ voucherID, chainID });
	const getChain = useChainState((s) => s.getChain);
	const [page, setPage] = useState({ offset: 0, size: 10 });
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

	const mergeTokens = (arrayA: any[], arrayB: any[]) => {
		const map = new Map();

		arrayA.forEach((token: any) => map.set(token.address, { ...token }));

		arrayB.forEach((token: any) => {
			const key = token.address || token.nft_contract_address;
			if (map.has(key)) {
				map.set(key, { ...map.get(key), ...token });
			}
		});

		return Array.from(map.values());
	};

	useEffect(() => {
		if (listTokens.data?.tokens.length > 0) {
			const tokens = mergeTokens(listTokens.data.tokens, infoTokens.data || []);
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
				<Table
					label='Voucher Token Holders'
					className='max-md:min-w-[1000px]'
					data={mergedTokens}
					columns={[
						{
							label: 'Token No',
							render: (d) => '#' + d.voucher_token_id,
						},
						{
							label: 'Credit name',
							render: (d) => (
								<Link className='text-[#0091C2] inline-block w-[300px] truncate text-center' href={`${voucherID}/credits/${d.address}/tokenNo/${d.voucher_token_id}`}>
									{d.project_name}
								</Link>
							),
						},
						chainID === CHAIN_ID_OFFSET
							? {
									label: 'Country',
									render: (d) => (
										<div className='flex items-center justify-center'>
											<img src={d.project_nation_flag} alt='Nation Flag' className='w-6 h-4' style={{ objectFit: 'cover' }} />
										</div>
									),
								}
							: null,
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
			</div>
		</div>
	);
};
