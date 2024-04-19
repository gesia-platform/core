const NotaryPublic = artifacts.require('NotaryPublic');

require('dotenv').config();

contract('Notary Public Registration Test', (accounts) => {
	let offsetNotaryPublic;

	const OFFSET_NOTARY_PUBLIC = process.env.OFFSET_NOTARY_PUBLIC;
	// const OFFSET_BLS_PUBLIC_KEY = process.env.OFFSET_BLS_PUBLIC_KEY;

	before(async () => {
		offsetNotaryPublic = await NotaryPublic.at(OFFSET_NOTARY_PUBLIC);
	});

	it('should grant and verify membership status', async () => {
		await offsetNotaryPublic.setMembership(process.env.EMISSION_IOA_ACCOUNT, true, { from: accounts[0] });
		const offsetIsMembership = await offsetNotaryPublic.hasMembership(accounts[0]);

		assert.equal(offsetIsMembership, true, 'The account should have membership status granted');
	});

	/**
	 * @dev 수동 테스트로 변경
	 */
	// it('should correctly register and retrieve a BLS public key', async () => {
	// 	await offsetNotaryPublic.register(OFFSET_BLS_PUBLIC_KEY, { from: accounts[0] });
	// 	const pubKey = await offsetNotaryPublic.getPubkey(accounts[0]);

	// 	assert.equal(pubKey, OFFSET_BLS_PUBLIC_KEY, 'The retrieved BLS public key should match the registered key');
	// });
});
