'use client';

import { useState, useEffect } from 'react';
import Image from 'next/image';
import useChainState from '@/stores/use-chain-state';
import { useGetNftsBalance } from '@/hooks/use-get-nfts-balance';

export const NFTList = ({ borderHidden, userAddress }: { borderHidden?: boolean; userAddress: string }) => {
	const { id: chainID } = useChainState();
	const { data: nftBalances, error } = useGetNftsBalance({ chainID, userAddress });

	if (error) {
		return <p>Error loading NFT data</p>;
	}

	const hasNfts = nftBalances?.some((nft: any) => nft.balance > 0);

	return (
		<div className={`mr-2 relative flex flex-col rounded-[8px] ${!borderHidden && 'border border-[#EAECED]'}`}>
			{!hasNfts ? (
				<p className='p-4 text-gray-500'>No NFTs owned.</p>
			) : (
				<table className='min-w-full divide-y divide-gray-200'>
					<thead>
						<tr>
							<th className='px-6 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider'>Carbon Offset NFT Address</th>
							<th className='px-6 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider'>Token ID</th>
							<th className='px-6 py-3 bg-gray-50 text-left text-xs font-medium text-gray-500 uppercase tracking-wider'>Balance</th>
						</tr>
					</thead>
					<tbody className='bg-white divide-y divide-gray-200'>
						{nftBalances
							?.filter((nft: any) => nft.balance > 0)
							.map((nft: any, index: number) => (
								<tr key={index}>
									<td className='px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900'>{nft.contract}</td>
									<td className='px-6 py-4 whitespace-nowrap text-sm text-gray-500'>{nft.tokenId}</td>
									<td className='px-6 py-4 whitespace-nowrap text-sm text-gray-500'>{nft.balance}</td>
								</tr>
							))}
					</tbody>
				</table>
			)}
		</div>
	);
};
