'use client';

import { ChainSelect } from '@/components/chain-select';
import { Pagination } from '@/components/pagination';
import { Table } from '@/components/table';
import { useListTokens } from '@/hooks/use-list-tokens';
import { useInfoTokens } from '@/hooks/use-info-tokens';
import useChainState from '@/stores/use-chain-state';
import Link from 'next/link';
import { useState, useEffect } from 'react';

export const TokenList = ({ type }: { type: 'ft' | 'nft' }) => {
	const chainID = useChainState((s) => s.id);
	const [page, setPage] = useState({ offset: 0, size: 10 });
	const [mergedTokens, setMergedTokens] = useState<any[]>([]);

	const listTokens = useListTokens({
		chainID,
		type,
		pageOffset: page.offset,
		pageSize: page.size,
	});
	const infoTokens = useInfoTokens({ type, chainID });

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

	return (
		<div className={`mt-5 ${type === 'nft' ? 'text-sm' : ''}`}>
			<Table
				className='max-md:min-w-[1000px]'
				label={`Total of ${BigInt(listTokens.data?.totalSize ?? 0).toLocaleString()} tokens`}
				headerComponent={
					<>
						<ChainSelect worker />
						<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={listTokens.data?.totalSize ?? 0} />
					</>
				}
				data={mergedTokens.length > 0 ? mergedTokens : listTokens.data?.tokens ?? []}
				columns={[
					type === 'ft'
						? {
								label: 'Token',
								render: (d) => d.symbol,
							}
						: null,
					type === 'ft'
						? {
								label: 'Type',
								render: (d) => d.type,
							}
						: null,
					type === 'nft'
						? {
								label: 'Carbon Project Name',
								render: (d) => (
									<span className='block truncate max-w-[300px]' title={d.project_name}>
										{d.project_name}
									</span>
								),
							}
						: null,
					type === 'nft'
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
						label: 'Contract Address',
						render: (d) => (
							<Link className='text-[#0091C2]' href={'/accounts/' + d.address}>
								{d.address}
							</Link>
						),
					},
					type === 'nft'
						? {
								label: 'Voucher Deposit',
								render: (d) =>
									d.deposit ? (
										<Link className='text-[#0091C2]' href={'/accounts/' + d.deposit}>
											{d.deposit}
										</Link>
									) : null,
							}
						: null,
					{
						label: 'Total Balance',
						render: (d) => BigInt(d.totalSupply).toLocaleString(),
					},

					{
						label: 'Holders',
						render: (d) => d.holders,
					},
				]}
				footerRightComponent={<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={listTokens.data?.totalSize ?? 0} />}
			/>
		</div>
	);
};
