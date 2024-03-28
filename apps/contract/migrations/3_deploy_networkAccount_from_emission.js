const NetworkAccount = artifacts.require('NetworkAccount');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(NetworkAccount, 'emission', 'http://3.39.139.167:8645/');
	const networkAccount = await NetworkAccount.deployed();

	console.log('NetworkAccount : ', networkAccount.address);
};
