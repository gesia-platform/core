'use client';

import { AddressTag } from '@/components/address-tag';
import { CarbonLabel } from '@/components/carbon-label';
import { Details } from '@/components/details';
import { DetailsRow } from '@/components/details-row';
import { DetailsRows } from '@/components/details-rows';
import { NFTList } from '@/components/nft-select';
import { CHAIN_ID_OFFSET } from '@/constants/chain';
import { useGetWeb3Account } from '@/hooks/use-get-web3-account';
import useChainState from '@/stores/use-chain-state';
import { formatNZC } from '@/utils/formatter';
import { redirect } from 'next/navigation';
import { useMemo } from 'react';

export const AccountDetails = ({ accountID }: { accountID: string }) => {
	const { id: chainID, getChain } = useChainState();
	const getWeb3Account = useGetWeb3Account({ accountID, chainID });

	const account = useMemo(() => {
		return getWeb3Account.data?.account || null;
	}, [getWeb3Account]);

	if (getWeb3Account.error) return redirect('/error');
	if (!account) return null;

	return (
		<div className='mt-5'>
			<Details grid headerComponent={<CarbonLabel chainID={chainID} />}>
				<DetailsRows>
					<DetailsRow label='Chain ID'>{`#${chainID} ${getChain()?.label}`}</DetailsRow>
					<DetailsRow label='Address'>
						{account.address}
						<AddressTag address={account.address} />
					</DetailsRow>

					<DetailsRow label='Balance'>{formatNZC(BigInt(account.balance))}</DetailsRow>
				</DetailsRows>
				{chainID === CHAIN_ID_OFFSET ? <NFTList borderHidden={false} userAddress={account.address} /> : null}
			</Details>
		</div>
	);
};
