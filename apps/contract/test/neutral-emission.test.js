const AppStore = artifacts.require('AppStore');
const AppPermission = artifacts.require('AppPermission');
const NetworkAccount = artifacts.require('NetworkAccount');

require('dotenv').config();

contract('AppStore and AppPermission Test', (accounts) => {
	let appStore;
	let appPermission;
	let networkAccount;
	const domain = 'https://test2-api.carbonmonster.kr/';
	const ip = '3.39.139.167';
	let appId = 1; // ! +1

	const appStoreAddress = '0xB0A568ac6a54F542C43882ea2E14a4bB3688b47a';
	const appPermissionAddress = '0x30947A48c5B64651e38DEA64cAdE0AeA1946B14b';
	const networkAccountAddress = '0x2D5c5950bB81379c969090733ee9c5970Eb3004d';

	before(async () => {
		appStore = await AppStore.at(appStoreAddress);
		appPermission = await AppPermission.at(appPermissionAddress);
		networkAccount = await NetworkAccount.at(networkAccountAddress);
	});

	it('should create a new app in the AppStore', async () => {
		const tx = await appStore.createApp(domain, accounts[0], { from: accounts[0] });
		const appCreatedEvent = tx.logs.find((e) => e.event === 'AppCreated');

		const { appID, domain: eventDomain, owner } = appCreatedEvent.args;

		assert.equal(appID, appId, "The appID in the event should match the created app's id");
		assert.equal(eventDomain, domain, "The domain in the event should match the created app's domain");
		assert.equal(owner, accounts[0], "The owner in the event should match the creator's address");
	});

	it('should allow a valid network access request via AppPermission', async () => {
		const tx = await appPermission.requestNetworkAccess(appId, networkAccount.address, ip, { from: accounts[0] });
		const networkAccessPermissionRequestedEvent = tx.logs.find((e) => e.event === 'NetworkAccessPermissionRequested');

		const { appID, networkAccount: eventNetworkAccount, ip: eventIp } = networkAccessPermissionRequestedEvent.args;

		assert.equal(appID.toNumber(), appId, 'App ID should match');
		assert.equal(eventNetworkAccount, networkAccount.address, 'Network account should match');
		assert.equal(eventIp, ip, 'IP should match');
	});
});
