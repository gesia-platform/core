const AppStore = artifacts.require('AppStore');
const AppPermission = artifacts.require('AppPermission');
const NetworkAccount = artifacts.require('NetworkAccount');

require('dotenv').config();

contract('AppStore and AppPermission Test', (accounts) => {
	let appStore;
	let appPermission;
	let networkAccount;
	let appId;

	const APP_STORE_ADDRESS = process.env.APP_STORE_ADDRESS;
	const APP_PERMISSION_ADDRESS = process.env.APP_PERMISSION_ADDRESS;
	const EMISSION_NETWORK_ACCOUNT_ADDRESS = process.env.EMISSION_NETWORK_ACCOUNT_ADDRESS;
	const APPLICATION_DOMAIN = process.env.APPLICATION_DOMAIN;
	const APPLICATION_IP = process.env.APPLICATION_IP;

	before(async () => {
		appStore = await AppStore.at(APP_STORE_ADDRESS);
		appPermission = await AppPermission.at(APP_PERMISSION_ADDRESS);
		networkAccount = await NetworkAccount.at(EMISSION_NETWORK_ACCOUNT_ADDRESS);
	});

	it('should create a new app in the AppStore', async () => {
		const tx = await appStore.createApp(APPLICATION_DOMAIN, accounts[0], { from: accounts[0] });
		const appCreatedEvent = tx.logs.find((e) => e.event === 'AppCreated');

		const { appID, domain: eventDomain, owner } = appCreatedEvent.args;
		appId = appID;

		assert.equal(eventDomain, APPLICATION_DOMAIN, "The domain in the event should match the created app's domain");
		assert.equal(owner, accounts[0], "The owner in the event should match the creator's address");
	});

	it('should allow a valid network access request via AppPermission', async () => {
		const tx = await appPermission.requestNetworkAccess(appId, networkAccount.address, web3.utils.asciiToHex(APPLICATION_IP), { from: accounts[0] });
		const networkAccessPermissionRequestedEvent = tx.logs.find((e) => e.event === 'NetworkAccessPermissionRequested');

		const { appID, networkAccount: eventNetworkAccount, ip: eventIp } = networkAccessPermissionRequestedEvent.args;

		assert.equal(appID.toNumber(), appId, 'App ID should match');
		assert.equal(eventNetworkAccount, networkAccount.address, 'Network account should match');
		assert.equal(web3.utils.hexToAscii(eventIp).replace(/\u0000/g, ''), APPLICATION_IP, 'IP should match');
	});
});
