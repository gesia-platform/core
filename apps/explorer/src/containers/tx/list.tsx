'use client';

import { ChainSelect } from '@/components/chain-select';
import { Pagination } from '@/components/pagination';
import { Spinner } from '@/components/spinner';
import { Table } from '@/components/table';
import { ToAddress } from '@/components/to-address';
import { useListTxs } from '@/hooks/use-list-txs';
import useChainState from '@/stores/use-chain-state';
import { formatAddress, formatNZC, formatHash, formatTimestampFromNow } from '@/utils/formatter';
import Link from 'next/link';
import { useState } from 'react';
import { SyncLoader } from 'react-spinners';

export const TxList = ({ blockID }: { blockID?: string }) => {
	const chainID = useChainState((s) => s.id);
	const [page, setPage] = useState({ offset: 0, size: 10 });

	const listTxs = useListTxs({
		chainID,
		blockID: blockID || undefined,
		pageOffset: page.offset,
		pageSize: page.size,
	});

	const isLoading = !listTxs.data;

	return (
		<div className='mt-5'>
			{isLoading ? (
				<Spinner color='#0091C2' size={8} />
			) : (
				<Table
					className='max-md:min-w-[1000px]'
					label={`More than ${BigInt(listTxs.data?.totalSize ?? 0).toLocaleString()} transactions found`}
					headerComponent={
						<>
							<ChainSelect />
							<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={listTxs.data?.totalSize ?? 0} />
						</>
					}
					data={listTxs.data?.txs ?? []}
					columns={[
						{
							label: 'Txn Hash',
							render: (d) => (
								<Link className='text-[#0091C2]' href={'/txs/' + d.hash}>
									{formatHash(d.hash)}
								</Link>
							),
						},
						{
							label: 'Block',
							render: (d) => (
								<Link className='text-[#0091C2]' href={'/blocks/' + d.block?.height}>
									{d.block?.height}
								</Link>
							),
						},
						{
							label: 'Age',
							render: (d) => formatTimestampFromNow(d.block?.timestamp),
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
							render: (d) => <ToAddress tx={d} label />,
						},
						{
							label: 'Value',
							render: (d) => formatNZC(BigInt(d.value)),
						},
						{
							label: 'Txn Fee',
							render: (d) => formatNZC(BigInt(d.effectiveGasPrice) * BigInt(d.gasUsed)),
						},
					]}
					footerRightComponent={<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={listTxs.data?.totalSize ?? 0} />}
				/>
			)}
		</div>
	);
};
