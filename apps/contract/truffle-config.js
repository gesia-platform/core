const HDWalletProvider = require('@truffle/hdwallet-provider');
require('dotenv').config();

module.exports = {
	networks: {
		gesia: {
			host: '52.78.12.174',
			port: 8545,
			network_id: '8888',
			provider: new HDWalletProvider(process.env.PRIVATE_KEY, 'http://52.78.12.174:8545/'),
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
