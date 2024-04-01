const AppPermission = artifacts.require('AppPermission');

require('dotenv').config();

contract('Network Access Permission Test', (accounts) => {
	let appPermission;

	const appPermissionAddress = process.env.APP_PERMISSION_ADDRESS;
	const networkAccountAddress = process.env.NETWORK_ACCOUNT_ADDRESS;

	before(async () => {
		appPermission = await AppPermission.at(appPermissionAddress);
	});

	it('should respond to a network access request for an app correctly', async () => {
		const tx = await appPermission.responseNetworkAccess(1, networkAccountAddress, '0x000000', false, { from: accounts[0] });
		const responseNetworkAccessEvent = tx.logs.find((e) => e.event === 'NetworkAccessPermissionResponsed');

		const { appID, networkAccount, signature, isGranted } = responseNetworkAccessEvent.args;

		assert.equal(appID, 1, "The appID in the event should match the requested app's ID");
		assert.equal(networkAccount, networkAccountAddress, "The networkAccount in the event should match the requested network account's address");
		assert.equal(signature, '0x000000', 'The signature in the event should match the provided signature');
		assert.equal(isGranted, false, 'The isGranted flag in the event should reflect the permission response correctly');
	});
});
