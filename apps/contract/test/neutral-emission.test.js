const AppStore = artifacts.require('AppStore');
const AppPermission = artifacts.require('AppPermission');
const NetworkAccount = artifacts.require('NetworkAccount');

require('dotenv').config();

contract('AppStore and AppPermission Test', (accounts) => {
	let appStore;
	let appPermission;
	let networkAccount;
	const domain = process.env.APPLICATION_DOMAIN;
	const ip = process.env.APPLICATION_IP;
	let appId;

	const appStoreAddress = process.env.APP_STORE_ADDRESS;
	const appPermissionAddress = process.env.APP_PERMISSION_ADDRESS;
	const networkAccountAddress = process.env.NETWORK_ACCOUNT_ADDRESS;

	before(async () => {
		appStore = await AppStore.at(appStoreAddress);
		appPermission = await AppPermission.at(appPermissionAddress);
		networkAccount = await NetworkAccount.at(networkAccountAddress);
	});

	it('should create a new app in the AppStore', async () => {
		const tx = await appStore.createApp(domain, accounts[0], { from: accounts[0] });
		const appCreatedEvent = tx.logs.find((e) => e.event === 'AppCreated');

		const { appID, domain: eventDomain, owner } = appCreatedEvent.args;
		appId = appID;

		assert.equal(eventDomain, domain, "The domain in the event should match the created app's domain");
		assert.equal(owner, accounts[0], "The owner in the event should match the creator's address");
	});

	it('should allow a valid network access request via AppPermission', async () => {
		const tx = await appPermission.requestNetworkAccess(appId, networkAccount.address, web3.utils.asciiToHex(ip), { from: accounts[0] });
		const networkAccessPermissionRequestedEvent = tx.logs.find((e) => e.event === 'NetworkAccessPermissionRequested');

		const { appID, networkAccount: eventNetworkAccount, ip: eventIp } = networkAccessPermissionRequestedEvent.args;

		assert.equal(appID.toNumber(), appId, 'App ID should match');
		assert.equal(eventNetworkAccount, networkAccount.address, 'Network account should match');
		assert.equal(web3.utils.hexToAscii(eventIp).replace(/\u0000/g, ''), ip, 'IP should match');
	});
});
