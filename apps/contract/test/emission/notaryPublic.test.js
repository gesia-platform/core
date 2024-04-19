const NotaryPublic = artifacts.require('NotaryPublic');

require('dotenv').config();

contract('Notary Public Registration Test', (accounts) => {
	let emissionNotaryPublic;

	const EMISSION_NOTARY_PUBLIC = process.env.EMISSION_NOTARY_PUBLIC;
	const EMISSION_BLS_PUBLIC_KEY = process.env.EMISSION_BLS_PUBLIC_KEY;

	before(async () => {
		emissionNotaryPublic = await NotaryPublic.at(EMISSION_NOTARY_PUBLIC);
	});

	it('should grant and verify membership status', async () => {
		await emissionNotaryPublic.setMembership(process.env.EMISSION_IOA_ACCOUNT, true, { from: accounts[0] });
		const emissionIsMembership = await emissionNotaryPublic.hasMembership(accounts[0]);

		assert.equal(emissionIsMembership, true, 'The account should have membership status granted');
	});

	/**
	 * @dev 수동 테스트로 변경
	 */
	// it('should correctly register and retrieve a BLS public key', async () => {
	// 	await emissionNotaryPublic.register(EMISSION_BLS_PUBLIC_KEY, { from: accounts[0] });
	// 	const pubKey = await emissionNotaryPublic.getPubkey(accounts[0]);

	// 	assert.equal(pubKey, EMISSION_BLS_PUBLIC_KEY, 'The retrieved BLS public key should match the registered key');
	// });
});
