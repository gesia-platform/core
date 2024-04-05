const CarbonEmissions = artifacts.require('CarbonEmissions');
const TestAmericanoCarbonEmissionsCalculator = artifacts.require('TestAmericanoCarbonEmissionsCalculator');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(CarbonEmissions, 'TestCarbonEmissions');
	const testCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(TestAmericanoCarbonEmissionsCalculator, testCarbonEmissions.address, process.env.NOTARY_PUBLIC_ADDRESS);
	const testAmericanoCarbonEmissionsCalculator = await TestAmericanoCarbonEmissionsCalculator.deployed();

	console.log('TestAmericanoCarbonEmissionsCalculator : ', testAmericanoCarbonEmissionsCalculator.address);
};
