import Web3 from 'web3';

export const CHAIN_ID_NEUTRALITY = 1;
export const CHAIN_ID_EMISSION = 2;
export const CHAIN_ID_OFFSET = 3;

export const CHAIN_NAME_NEUTRALITY = 'neutrality';
export const CHAIN_NAME_EMISSION = 'emission';
export const CHAIN_NAME_OFFSET = 'offset';

export const CHAIN_LABEL_NEUTRALITY = 'Net zero sequence chain';
export const CHAIN_LABEL_EMISSION = 'Carbon emissions tracking chain';
export const CHAIN_LABEL_OFFSET = 'Carbon credit chain';

export const CHAINS = [
	{
		id: CHAIN_ID_NEUTRALITY,
		name: CHAIN_NAME_NEUTRALITY,
		label: CHAIN_LABEL_NEUTRALITY,
	},
	{
		id: CHAIN_ID_EMISSION,
		name: CHAIN_NAME_EMISSION,
		label: CHAIN_LABEL_EMISSION,
	},
	{ id: CHAIN_ID_OFFSET, name: CHAIN_NAME_OFFSET, label: CHAIN_LABEL_OFFSET },
];

export const CHAIN_VOUCHER_LOG_PARAM_TYPES = ['address', 'address', 'address', 'uint256', 'uint256'];

export const CHAIN_VOUCHER_EVENT = 'event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)';

export const CHAIN_VOUCHER_LOG_TOPIC_0_HASH = Web3.utils.sha3('TransferSingle(address,address,address,uint256,uint256)');
