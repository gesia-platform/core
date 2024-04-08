const CarbonEmissions = artifacts.require('CarbonEmissions');
const TestAmericanoCarbonEmissionsCalculator = artifacts.require('TestAmericanoCarbonEmissionsCalculator');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(CarbonEmissions, 'TestCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const testCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(TestAmericanoCarbonEmissionsCalculator, testCarbonEmissions.address, process.env.EMISSION_NOTARY_PUBLIC);
	const testAmericanoCarbonEmissionsCalculator = await TestAmericanoCarbonEmissionsCalculator.deployed();
	await testCarbonEmissions.setCalculatorApproval(testAmericanoCarbonEmissionsCalculator.address, true);

	console.log('TestAmericanoCarbonEmissionsCalculator : ', testAmericanoCarbonEmissionsCalculator.address);
};
