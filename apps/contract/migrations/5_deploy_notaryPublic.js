const NotaryPublic = artifacts.require('NotaryPublic');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(NotaryPublic);
	const notaryPublic = await NotaryPublic.deployed();

	console.log('NotaryPublic : ', notaryPublic.address);
};
