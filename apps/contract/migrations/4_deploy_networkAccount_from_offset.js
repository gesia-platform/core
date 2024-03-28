const NetworkAccount = artifacts.require('NetworkAccount');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(NetworkAccount, 'offset', 'http://43.200.218.66:8445/');
	const networkAccount = await NetworkAccount.deployed();

	console.log('NetworkAccount : ', networkAccount.address);
};
