import { Main } from '@/components/main';
import { PageLabel } from '@/components/page-label';
import { TokenList } from '@/containers/token/list';

export default function NFTs() {
	return (
		<Main>
			<PageLabel>Carbon Credits List</PageLabel>
			<TokenList type='nft' />
		</Main>
	);
}
