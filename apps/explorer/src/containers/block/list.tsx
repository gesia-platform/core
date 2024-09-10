'use client';

import { ChainSelect } from '@/components/chain-select';
import { Pagination } from '@/components/pagination';
import { Table } from '@/components/table';
import { Spinner } from '@/components/spinner';
import { CHAIN_ID_NEUTRALITY } from '@/constants/chain';
import { useListBlocks } from '@/hooks/use-list-blocks';
import useChainState from '@/stores/use-chain-state';
import { formatAddress, formatNZC, formatTimestampFromNow } from '@/utils/formatter';
import Link from 'next/link';
import { useState } from 'react';

export const BlockList = ({}) => {
	const chainID = useChainState((s) => s.id);
	const [page, setPage] = useState({ offset: 0, size: 10 });

	const { data: listBlocks, error } = useListBlocks({
		chainID,
		pageOffset: page.offset,
		pageSize: page.size,
	});

	const isLoading = !listBlocks && !error;

	return (
		<div className='mt-5'>
			{isLoading ? (
				<Spinner size={8} />
			) : (
				<Table
					className='max-md:min-w-[1000px]'
					label={`Total of ${BigInt(listBlocks?.totalSize ?? 0).toLocaleString()} blocks`}
					headerComponent={
						<>
							<ChainSelect />
							<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={listBlocks?.totalSize ?? 0} />
						</>
					}
					data={listBlocks?.blocks ?? []}
					columns={[
						{
							label: 'Block',
							render: (d) => (
								<Link className='text-[#0091C2]' href={'/blocks/' + d.height}>
									{d.height}
								</Link>
							),
						},
						{ label: 'Age', render: (d) => formatTimestampFromNow(d.timestamp) },
						{
							label: 'Txn',
							render: (d) => (
								<Link className='text-[#0091C2]' href={'/txs?block=' + d.height}>
									{d.txns}
								</Link>
							),
						},
						{
							label: chainID === CHAIN_ID_NEUTRALITY ? 'Validator' : 'Signer',
							render: (d) => (
								<Link className='text-[#0091C2]' href={'/accounts/' + d.miner}>
									{formatAddress(d.miner)}
								</Link>
							),
						},
						{
							label: 'Gas Used',
							render: (d) => BigInt(d.gasUsed).toLocaleString(),
						},
						{
							label: 'Gas Limit',
							render: (d) => BigInt(d.gasLimit).toLocaleString(),
						},
						{
							label: 'Reward',
							render: (d) => {
								return formatNZC(d.txs?.reduce((p: any, c: any) => p + BigInt(c.effectiveGasPrice) * BigInt(c.gasUsed), BigInt(0)));
							},
						},
					]}
					footerRightComponent={<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={listBlocks?.totalSize ?? 0} />}
				/>
			)}
		</div>
	);
};
