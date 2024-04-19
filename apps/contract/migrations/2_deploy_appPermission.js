const AppStore = artifacts.require('AppStore');
const AppPermission = artifacts.require('AppPermission');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(AppStore);
	const appStore = await AppStore.deployed();
	await deployer.deploy(AppPermission, appStore.address);
	const appPermission = await AppPermission.deployed();

	console.log('AppStore : ', appStore.address);
	console.log('AppPermission : ', appPermission.address);
};
