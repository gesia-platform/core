const HDWalletProvider = require('@truffle/hdwallet-provider');
require('dotenv').config();

module.exports = {
	networks: {
		development: {
			host: '52.78.12.174',
			port: 8545,
			network_id: '8888',
			provider: new HDWalletProvider(process.env.PRIVATE_KEY, 'http://52.78.12.174:8545/'),
		},
		emission: {
			host: '3.39.139.167',
			port: 8645,
			network_id: '5555',
			provider: new HDWalletProvider(process.env.EMISSION_PRIVATE_KEY, 'http://3.39.139.167:8645/'),
		},
		neutral: {
			host: '13.209.79.57',
			port: 8545,
			network_id: '5555',
			provider: new HDWalletProvider(process.env.NEUTRAL_PRIVATE_KEY, 'http://13.209.79.57:8545/'),
		},
		neutral_emission: {
			host: '13.209.79.57',
			port: 8545,
			network_id: '5555',
			provider: new HDWalletProvider(process.env.EMISSION_PRIVATE_KEY, 'http://13.209.79.57:8545/'),
		},
		neutral_offset: {
			host: '13.209.79.57',
			port: 8545,
			network_id: '5555',
			provider: new HDWalletProvider(process.env.OFFSET_PRIVATE_KEY, 'http://13.209.79.57:8545/'),
		},
		offset: {
			host: '43.200.218.66',
			port: 8445,
			network_id: '5555',
			provider: new HDWalletProvider(process.env.OFFSET_PRIVATE_KEY, 'http://43.200.218.66:8445/'),
		},
	},
	compilers: {
		solc: {
			version: '0.8.19',
			settings: {
				optimizer: {
					enabled: true,
					runs: 1000,
				},
			},
		},
	},
};
