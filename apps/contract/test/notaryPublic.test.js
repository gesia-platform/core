const NotaryPublic = artifacts.require('NotaryPublic');

require('dotenv').config();

contract('Notary Public Registration Test', (accounts) => {
	let notaryPublic;
	const blsPublicKey = process.env.BLS_PUBLIC_KEY;
	const notaryPublicAddress = process.env.NOTARY_PUBLIC;

	before(async () => {
		notaryPublic = await NotaryPublic.at(notaryPublicAddress);
	});

	it('should correctly register and retrieve a BLS public key', async () => {
		await notaryPublic.register(blsPublicKey, { from: accounts[0] });
		const pubKey = await notaryPublic.getPubkey(accounts[0]);

		assert.equal(pubKey, blsPublicKey, 'The retrieved BLS public key should match the registered key');
	});
});
