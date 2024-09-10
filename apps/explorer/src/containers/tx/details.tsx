'use client';

import { AddressTag } from '@/components/address-tag';
import { CarbonLabel } from '@/components/carbon-label';
import { Details } from '@/components/details';
import { DetailsRow } from '@/components/details-row';
import { DetailsRows } from '@/components/details-rows';
import { ToAddress } from '@/components/to-address';
import { TxInputData } from '@/components/tx-input-data';
import { CHAIN_ID_EMISSION, CHAIN_VOUCHER_LOG_TOPIC_0_HASH } from '@/constants/chain';
import { useGetTx } from '@/hooks/use-get-tx';
import { useGetWeb3Account } from '@/hooks/use-get-web3-account';
import useChainState from '@/stores/use-chain-state';
import { formatNZC, formatTCO2, formatTimestamp, formatTimestampFromNow } from '@/utils/formatter';
import Link from 'next/link';
import { redirect } from 'next/navigation';
import { useMemo, useRef } from 'react';
import Web3 from 'web3';

export const TxDetails = ({ txID }: { txID: string }) => {
	const chainID = useChainState((s) => s.id);
	const getTx = useGetTx({ txID, chainID });
	const getChain = useChainState((s) => s.getChain);

	const tx = useMemo(() => {
		return getTx.data?.tx || null;
	}, [getTx.data]);

	const abi = useRef(new Web3().eth.abi).current;

	const chain = useMemo(() => {
		const chainID = tx?.block?.chainID;
		return chainID ? getChain(chainID) : null;
	}, [getChain, tx]);

	const getWeb3AccountTo = useGetWeb3Account(tx ? { accountID: tx.to, chainID } : undefined);

	const decodeLogs = (logs: any[]) => {
		const transferEventAbi = {
			anonymous: false,
			inputs: [
				{ indexed: true, internalType: 'address', name: 'operator', type: 'address' },
				{ indexed: true, internalType: 'address', name: 'from', type: 'address' },
				{ indexed: true, internalType: 'address', name: 'to', type: 'address' },
				{ indexed: false, internalType: 'uint256', name: 'id', type: 'uint256' },
				{ indexed: false, internalType: 'uint256', name: 'value', type: 'uint256' },
			],
			name: 'TransferSingle',
			type: 'event',
		};

		return logs
			.map((log) => {
				try {
					const eventSignature = log.topics[0];
					const encodedSignature = abi.encodeEventSignature(transferEventAbi);

					if (eventSignature !== encodedSignature) {
						console.error('Event signature does not match');
						return null;
					}

					const decodedLog = abi.decodeLog(transferEventAbi.inputs, log.data, log.topics.slice(1));
					return decodedLog;
				} catch (error) {
					console.error('Error decoding log:', error);
					return null;
				}
			})
			.filter((result) => result !== null);
	};

	const renderVoucher = () => {
		const logsWithVoucher = tx.logs?.filter((x: any) => x.topics?.includes(CHAIN_VOUCHER_LOG_TOPIC_0_HASH));

		if (!logsWithVoucher || logsWithVoucher.length === 0) return null;

		const decodedLogs = decodeLogs(logsWithVoucher);

		return (
			<DetailsRow label='Event'>
				<div className='space-y-4'>
					{decodedLogs.map((decodedLog: any, index) => {
						if (!decodedLog) return null;

						const stringifyBigInt = (obj: any) => {
							return JSON.stringify(obj, (key, value) => (typeof value === 'bigint' ? value.toString() : value));
						};

						const tokenID = Number(BigInt(decodedLog.id));
						const value = formatTCO2(BigInt(decodedLog.value).toString(), chainID === CHAIN_ID_EMISSION);
						const fromAddress: string = decodedLog.from;
						const toAddress = decodedLog.to;

						return (
							<div key={index} className='p-4 border border-gray-200 rounded-lg shadow-md'>
								<div className='mb-2'>
									<Link className='text-[#0091C2]' href={'/accounts/' + logsWithVoucher[index].address}>
										{logsWithVoucher[index].address}
									</Link>
									<span className='ml-2'>(Contract)</span>
								</div>
								<div className='grid grid-cols-1 gap-2 mt-4'>
									<span>Token(App) ID: {tokenID.toString()}</span>
									<span>
										Amount: {(Number(value) / 1000).toString()}
										{chainID === 3 ? ' tCOC' : ' tCO2'}
									</span>
									<div className='mt-2'>
										<div>
											From:{' '}
											<Link className='text-[#0091C2]' href={'/accounts/' + fromAddress}>
												{fromAddress}
											</Link>
										</div>
										<div>
											To:{' '}
											<Link className='text-[#0091C2]' href={'/accounts/' + toAddress}>
												{toAddress}
											</Link>
										</div>
									</div>
								</div>
							</div>
						);
					})}
				</div>
			</DetailsRow>
		);
	};

	if (getTx.error) return redirect('/error');

	if (!tx || !chain) return null;

	return (
		<div className='mt-5'>
			<Details grid headerComponent={<CarbonLabel chainID={chainID} />}>
				<DetailsRows>
					<DetailsRow label='Chain ID'>{`#${chain.id} ${chain.label}`}</DetailsRow>
					<DetailsRow label='Transaction Hash'>{tx.hash}</DetailsRow>
				</DetailsRows>

				<DetailsRows>
					<DetailsRow label='Status'>{tx.status === '0' ? 'Failure' : 'Success'}</DetailsRow>
					<DetailsRow label='Block'>
						<Link className='text-[#0091C2]' href={'/blocks/' + tx.block.height}>
							{tx.block.height}
						</Link>
					</DetailsRow>

					<DetailsRow label='Timestamp'>
						{formatTimestampFromNow(tx.block.timestamp)} ({formatTimestamp(tx.block.timestamp)})
					</DetailsRow>
				</DetailsRows>

				<DetailsRows>
					<DetailsRow label='From'>
						<Link className='text-[#0091C2]' href={'/accounts/' + tx.from}>
							{tx.from}
						</Link>
						<AddressTag address={tx.from} />
					</DetailsRow>
					<DetailsRow label='To'>
						<ToAddress tx={tx} tag />
					</DetailsRow>

					{getWeb3AccountTo.data?.account.isContract && renderVoucher()}
				</DetailsRows>

				<DetailsRows>
					<DetailsRow label='Value'>{formatNZC(tx.value)}</DetailsRow>
					<DetailsRow label='Transaction Fee'>{formatNZC(BigInt(tx.effectiveGasPrice) * BigInt(tx.gasUsed))}</DetailsRow>
				</DetailsRows>

				<DetailsRows>
					<DetailsRow label='Gas Price'>{formatNZC(BigInt(tx.gasPrice))}</DetailsRow>
					<DetailsRow label='Gas Limit'>{BigInt(tx.gas).toLocaleString()}</DetailsRow>
				</DetailsRows>

				<DetailsRows>
					<DetailsRow label='Input Data'>
						<TxInputData input={tx.input} />
					</DetailsRow>
				</DetailsRows>
			</Details>
		</div>
	);
};
