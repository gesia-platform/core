const NotaryPublic = artifacts.require('NotaryPublic');

require('dotenv').config();

contract('AppStore and AppPermission Test', (accounts) => {
	let notaryPublic;
	const blsPublicKey = process.env.BLS_PUBLIC_KEY;
	const notaryPublicAddress = process.env.NOTARY_PUBLIC;

	before(async () => {
		notaryPublic = await NotaryPublic.at(notaryPublicAddress);
	});

	it('should create a new app in the AppStore', async () => {
		await notaryPublic.register(blsPublicKey, { from: accounts[0] });
		const pubKey = await notaryPublic.getPubkey(accounts[0]);

		assert.equal(pubKey, blsPublicKey, "The domain in the event should match the created app's domain");
	});
});
