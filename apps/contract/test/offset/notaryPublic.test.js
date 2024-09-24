const NotaryPublic = artifacts.require('NotaryPublic');

require('dotenv').config();

contract('Notary Public Registration Test', (accounts) => {
	let offsetNotaryPublic;

	const OFFSET_NOTARY_PUBLIC = process.env.OFFSET_NOTARY_PUBLIC;
	const OFFSET_BLS_PUBLIC_KEY = process.env.OFFSET_BLS_PUBLIC_KEY;

	before(async () => {
		offsetNotaryPublic = await NotaryPublic.at(OFFSET_NOTARY_PUBLIC);
	});

	/**
	 * @dev 1번
	 */
	it('should grant and verify membership status', async () => {
		await offsetNotaryPublic.setMembership(process.env.OFFSET_IOA_ACCOUNT, true, { from: accounts[0] });
		const offsetIsMembership = await offsetNotaryPublic.hasMembership(accounts[0]);

		assert.equal(offsetIsMembership, true, 'The account should have membership status granted');
	});

	/**
	 * @dev 2번
	 */
	it('should correctly register and retrieve a BLS public key', async () => {
		await offsetNotaryPublic.register(OFFSET_BLS_PUBLIC_KEY, { from: accounts[0] });
		const pubKey = await offsetNotaryPublic.getPubkey(accounts[0]);

		assert.equal(pubKey, OFFSET_BLS_PUBLIC_KEY, 'The retrieved BLS public key should match the registered key');
	});

	/**
	 * @dev 3번 / AppID 생성 후 계정 추가할 때 사용
	 */
	it('notarize', async () => {
		const prefix = '0x01'; // Use a valid 1-byte hex string for prefix
		const appID = 1;
		const signature = '0xb1931be82f6a900079ff15f2e8bed9c62aaa787884523ab1f8c02538ecfb27c9940025dd8f7068c864c0991875f9f0020504525935a0fc6dde271a42f6bd61ce225f5755043bb0d5533291fc6b799e84289bdcf86d310a89aab40b77c760e817';
		await offsetNotaryPublic.notarize(prefix, appID, signature, { from: accounts[0] });
		console.log('Successfully notarized');
	});
});
