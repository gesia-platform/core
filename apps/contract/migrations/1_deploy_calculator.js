const CarbonEmissions = artifacts.require('CarbonEmissions');
const BeefCarbonEmissionsCalculator = artifacts.require('BeefCarbonEmissionsCalculator');
const BicycleCarbonEmissionsCalculator = artifacts.require('BicycleCarbonEmissionsCalculator');
const CannedTunaCarbonEmissionsCalculator = artifacts.require('CannedTunaCarbonEmissionsCalculator');
const ChickenCarbonEmissionsCalculator = artifacts.require('ChickenCarbonEmissionsCalculator');
const EggCarbonEmissionsCalculator = artifacts.require('EggCarbonEmissionsCalculator');
const EmailCarbonEmissionsCalculator = artifacts.require('EmailCarbonEmissionsCalculator');
const LambCarbonEmissionsCalculator = artifacts.require('LambCarbonEmissionsCalculator');
const MilkCarbonEmissionsCalculator = artifacts.require('MilkCarbonEmissionsCalculator');
const MobileDataCarbonEmissionsCalculator = artifacts.require('MobileDataCarbonEmissionsCalculator');
const PhoneCallCarbonEmissionsCalculator = artifacts.require('PhoneCallCarbonEmissionsCalculator');
const PorkCarbonEmissionsCalculator = artifacts.require('PorkCarbonEmissionsCalculator');
const SalmonCarbonEmissionsCalculator = artifacts.require('SalmonCarbonEmissionsCalculator');
const WalkCarbonEmissionsCalculator = artifacts.require('WalkCarbonEmissionsCalculator');
const YoutubeCarbonEmissionsCalculator = artifacts.require('YoutubeCarbonEmissionsCalculator');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(CarbonEmissions, 'BeefCarbonEmissions');
	const beefCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(BeefCarbonEmissionsCalculator, beefCarbonEmissions.address);
	const beefCarbonEmissionsCalculator = await BeefCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'BicycleCarbonEmissions');
	const bicycleCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(BicycleCarbonEmissionsCalculator, bicycleCarbonEmissions.address);
	const bicycleCarbonEmissionsCalculator = await BicycleCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'CannedTunaCarbonEmissions');
	const cannedTunaCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(CannedTunaCarbonEmissionsCalculator, cannedTunaCarbonEmissions.address);
	const cannedTunaCarbonEmissionsCalculator = await CannedTunaCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'CannedTunaCarbonEmissions');
	const chickenCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(ChickenCarbonEmissionsCalculator, chickenCarbonEmissions.address);
	const chickenCarbonEmissionsCalculator = await ChickenCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'EggCarbonEmissions');
	const eggCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(EggCarbonEmissionsCalculator, eggCarbonEmissions.address);
	const eggCarbonEmissionsCalculator = await EggCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'EmailCarbonEmissions');
	const emailCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(EmailCarbonEmissionsCalculator, emailCarbonEmissions.address);
	const emailCarbonEmissionsCalculator = await EmailCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'LambCarbonEmissions');
	const lambCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(LambCarbonEmissionsCalculator, lambCarbonEmissions.address);
	const lambCarbonEmissionsCalculator = await LambCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'MilkCarbonEmissions');
	const milkCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(MilkCarbonEmissionsCalculator, milkCarbonEmissions.address);
	const milkCarbonEmissionsCalculator = await MilkCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'MobileCarbonEmissions');
	const mobileCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(MobileDataCarbonEmissionsCalculator, mobileCarbonEmissions.address);
	const mobileDataCarbonEmissionsCalculator = await MobileDataCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'PhoneCarbonEmissions');
	const phoneCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(PhoneCallCarbonEmissionsCalculator, phoneCarbonEmissions.address);
	const phoneCallCarbonEmissionsCalculator = await PhoneCallCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'PorkCarbonEmissions');
	const porkCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(PorkCarbonEmissionsCalculator, porkCarbonEmissions.address);
	const porkCarbonEmissionsCalculator = await PorkCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'SalmonCarbonEmissions');
	const salmonCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(SalmonCarbonEmissionsCalculator, salmonCarbonEmissions.address);
	const salmonCarbonEmissionsCalculator = await SalmonCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'WalkCarbonEmissions');
	const walkCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(WalkCarbonEmissionsCalculator, walkCarbonEmissions.address);
	const walkCarbonEmissionsCalculator = await WalkCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'YoutubeCarbonEmissions');
	const youtubeCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(YoutubeCarbonEmissionsCalculator, youtubeCarbonEmissions.address);
	const youtubeCarbonEmissionsCalculator = await YoutubeCarbonEmissionsCalculator.deployed();

	console.log('BeefCarbonEmissions : ', beefCarbonEmissions.address);
	console.log('BeefCarbonEmissionsCalculator : ', beefCarbonEmissionsCalculator.address);
	console.log('BicycleCarbonEmissions : ', bicycleCarbonEmissions.address);
	console.log('BicycleCarbonEmissionsCalculator : ', bicycleCarbonEmissionsCalculator.address);
	console.log('CannedTunaCarbonEmissions : ', cannedTunaCarbonEmissions.address);
	console.log('CannedTunaCarbonEmissionsCalculator : ', cannedTunaCarbonEmissionsCalculator.address);
	console.log('ChickenCarbonEmissions : ', chickenCarbonEmissions.address);
	console.log('ChickenCarbonEmissionsCalculator : ', chickenCarbonEmissionsCalculator.address);
	console.log('EggCarbonEmissions : ', eggCarbonEmissions.address);
	console.log('EggCarbonEmissionsCalculator : ', eggCarbonEmissionsCalculator.address);
	console.log('EmailCarbonEmissions : ', emailCarbonEmissions.address);
	console.log('EmailCarbonEmissionsCalculator : ', emailCarbonEmissionsCalculator.address);
	console.log('LambCarbonEmissions : ', lambCarbonEmissions.address);
	console.log('LambCarbonEmissionsCalculator : ', lambCarbonEmissionsCalculator.address);
	console.log('MilkCarbonEmissions : ', milkCarbonEmissions.address);
	console.log('MilkCarbonEmissionsCalculator : ', milkCarbonEmissionsCalculator.address);
	console.log('MobileCarbonEmissions : ', mobileCarbonEmissions.address);
	console.log('MobileDataCarbonEmissionsCalculator : ', mobileDataCarbonEmissionsCalculator.address);
	console.log('PhoneCarbonEmissions : ', phoneCarbonEmissions.address);
	console.log('PhoneCallCarbonEmissionsCalculator : ', phoneCallCarbonEmissionsCalculator.address);
	console.log('PorkCarbonEmissions : ', porkCarbonEmissions.address);
	console.log('PorkCarbonEmissionsCalculator : ', porkCarbonEmissionsCalculator.address);
	console.log('SalmonCarbonEmissions : ', salmonCarbonEmissions.address);
	console.log('SalmonCarbonEmissionsCalculator : ', salmonCarbonEmissionsCalculator.address);
	console.log('WalkCarbonEmissions : ', walkCarbonEmissions.address);
	console.log('WalkCarbonEmissionsCalculator : ', walkCarbonEmissionsCalculator.address);
	console.log('YoutubeCarbonEmissions : ', youtubeCarbonEmissions.address);
	console.log('YoutubeCarbonEmissionsCalculator : ', youtubeCarbonEmissionsCalculator.address);
};
