const NetworkAccount = artifacts.require('NetworkAccount');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(NetworkAccount, 'offset', process.env.OFFSET_NETWORK);
	const networkAccount = await NetworkAccount.deployed();

	console.log('NetworkAccount : ', networkAccount.address);
};
