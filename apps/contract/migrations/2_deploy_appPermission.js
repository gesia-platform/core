const AppStore = artifacts.require('AppStore');
const AppPermission = artifacts.require('AppPermission');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(AppStore);
	const appstore = await AppStore.deployed();
	await deployer.deploy(AppPermission, appstore.address);
	const appPermission = await AppPermission.deployed();

	console.log('AppStore : ', appstore.address);
	console.log('AppPermission : ', appPermission.address);
};
