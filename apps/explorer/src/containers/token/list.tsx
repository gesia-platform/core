'use client';

import { ChainSelect } from '@/components/chain-select';
import { Pagination } from '@/components/pagination';
import { Table } from '@/components/table';
import { useListTokens } from '@/hooks/use-list-tokens';
import useChainState from '@/stores/use-chain-state';
import Link from 'next/link';
import { useState } from 'react';

export const TokenList = ({ type }: { type: 'ft' | 'nft' }) => {
	const chainID = useChainState((s) => s.id);
	const [page, setPage] = useState({ offset: 0, size: 10 });

	const listTokens = useListTokens({
		chainID,
		type,
		pageOffset: page.offset,
		pageSize: page.size,
	});

	return (
		<div className='mt-5'>
			<Table
				className='max-md:min-w-[1000px]'
				label={`Total of ${BigInt(listTokens.data?.totalSize ?? 0).toLocaleString()} tokens`}
				headerComponent={
					<>
						<ChainSelect worker />
						<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={listTokens.data?.totalSize ?? 0} />
					</>
				}
				data={listTokens.data?.tokens ?? []}
				columns={[
					{
						label: 'Token',
						render: (d) => d.symbol,
					},
					{
						label: 'Type',
						render: (d) => d.type,
					},
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
