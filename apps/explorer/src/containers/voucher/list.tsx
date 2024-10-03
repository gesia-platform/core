'use client';

import { ChainSelect } from '@/components/chain-select';
import { Pagination } from '@/components/pagination';
import { Table } from '@/components/table';
import { useListVouchers } from '@/hooks/use-list-vouchers';
import useChainState from '@/stores/use-chain-state';
import Link from 'next/link';
import { useState } from 'react';

export const VoucherList = ({}) => {
	const chainID = useChainState((s) => s.id);
	const [page, setPage] = useState({ offset: 0, size: 10 });

	const listVouchers = useListVouchers({
		chainID,
		pageOffset: page.offset,
		pageSize: page.size,
	});

	return (
		<div className='mt-5'>
			<Table
				className='max-md:min-w-[1000px]'
				label={`Total of ${BigInt(listVouchers.data?.totalSize ?? 0).toLocaleString()} vouchers`}
				headerComponent={
					<>
						<ChainSelect worker />
					</>
				}
				data={listVouchers.data?.vouchers ?? []}
				columns={[
					{
						label: 'No',
						render: (_, index) => page.offset + index + 1,
					},
					{
						label: 'Registry Name',
						render: (d) => (
							<Link className='text-[#0091C2] inline-block w-[250px] truncate text-center' href={'/vouchers/' + d.address}>
								{d.name}
							</Link>
						),
					},
					{
						label: 'Token Type',
						render: (d) => d.type,
					},
					{
						label: 'Credit Projects',
						render: (d) => d.count,
					},
					{
						label: 'Contract Address',
						render: (d) => (
							<Link className='text-[#0091C2] underline' href={'/accounts/' + d.address}>
								{d.address}
							</Link>
						),
					},
				]}
				footerRightComponent={<Pagination onOffsetChange={(offset) => setPage((prev) => ({ ...prev, offset }))} offset={page.offset} size={page.size} totalSize={listVouchers.data?.totalSize ?? 0} />}
			/>
		</div>
	);
};
