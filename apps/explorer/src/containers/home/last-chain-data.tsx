'use client';

import { LastBlockItem } from '@/components/last-block-item';
import { LastList } from '@/components/last-list';
import { LastTransactionItem } from '@/components/last-transaction-item';
import { useListBlocks } from '@/hooks/use-list-blocks';
import { useEffect } from 'react';
import { formatTimestampFromNow } from '@/utils/formatter';
import useNetworkStore from '@/stores/use-chain-state';
import { useListTxs } from '@/hooks/use-list-txs';
import { SyncLoader } from 'react-spinners';
import { Spinner } from '@/components/spinner';

export const LastChainData = ({}) => {
	const chainStore = useNetworkStore();

	const listBlocks = useListBlocks({
		chainID: chainStore.id,
		pageOffset: 0,
		pageSize: 5,
	});

	const listTxs = useListTxs({
		chainID: chainStore.id,
		pageOffset: 0,
		pageSize: 5,
	});

	useEffect(() => {
		let t = setInterval(() => {
			listBlocks.mutate();
			listTxs.mutate();
		}, 5000);

		return () => {
			clearTimeout(t);
		};
	}, [chainStore.id]);

	return (
		<div className='bg-white border border-[#eaeced] shadow-md rounded-[8px] mt-5'>
			<div className='grid grid-cols-3'>
				{chainStore.chains.map((data) => (
					<button
						key={data.name}
						onClick={() => {
							chainStore.setID(data.id);
						}}
						className={`h-[72px] max-md:!h-[60px] px-1 rounded-tl-[8px] font-medium text-[16px] ${chainStore.id === data.id ? 'text-[#00C1B3] bg-white border-r' : 'bg-[#f5f6f7]'}`}
					>
						Chain ID #{data.id} ({data.label})
					</button>
				))}
			</div>

			<div className='grid grid-cols-2 gap-x-5 max-lg:grid-cols-1 max-lg:gap-x-0 px-5 py-[30px] max-lg:!py-0'>
				<LastList label='Last Blocks' moreHref='/blocks' moreLabel='VIEW ALL BLOCKS'>
					{!listBlocks.data ? (
						<Spinner height='300px' color='#000000' />
					) : (
						listBlocks.data.blocks.map((block: any, index: number) => {
							return <LastBlockItem minedLabel={chainStore.getChain()?.consensusAlgorithm === 'PoS' ? 'Validated' : 'Signed'} key={chainStore.id + '' + index} number={block.height} txns={block.txns} time={formatTimestampFromNow(block.timestamp)} proposer={block.miner} />;
						})
					)}
				</LastList>

				<LastList label='Last Transactions' moreHref='/txs' moreLabel='VIEW ALL TRANSACTIONS'>
					{!listTxs.data ? <Spinner height='300px' color='#000000' /> : listTxs.data.txs.map((tx: any, i: number) => <LastTransactionItem key={chainStore.id + '' + i} hash={tx.hash} time={formatTimestampFromNow(tx.block?.timestamp)} from={tx.from} to={tx.to} />)}
				</LastList>
			</div>
		</div>
	);
};
