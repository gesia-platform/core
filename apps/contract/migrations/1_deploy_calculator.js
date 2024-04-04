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
const ElectricityCarbonEmissionsCalculator = artifacts.require('ElectricityCarbonEmissionsCalculator');
const WaterCarbonEmissionsCalculator = artifacts.require('WaterCarbonEmissionsCalculator');
const GasCarbonEmissionsCalculator = artifacts.require('GasCarbonEmissionsCalculator');
const PublicTransportCarbonEmissionsCalculator = artifacts.require('PublicTransportCarbonEmissionsCalculator');
const CarCarbonEmissionsCalculator = artifacts.require('CarCarbonEmissionsCalculator');
const AirplaneCarbonEmissionsCalculator = artifacts.require('AirplaneCarbonEmissionsCalculator');
const TrainCarbonEmissionsCalculator = artifacts.require('TrainCarbonEmissionsCalculator');
const AmericanoCarbonEmissionsCalculator = artifacts.require('AmericanoCarbonEmissionsCalculator');

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

	await deployer.deploy(CarbonEmissions, 'ElectricityCarbonEmissions');
	const electricityCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(ElectricityCarbonEmissionsCalculator, electricityCarbonEmissions.address);
	const electricityCarbonEmissionsCalculator = await ElectricityCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'WaterCarbonEmissions');
	const waterCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(WaterCarbonEmissionsCalculator, waterCarbonEmissions.address);
	const waterCarbonEmissionsCalculator = await WaterCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'GasCarbonEmissions');
	const gasCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(GasCarbonEmissionsCalculator, gasCarbonEmissions.address);
	const gasCarbonEmissionsCalculator = await GasCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'PublicTransportCarbonEmissions');
	const publicTransportCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(PublicTransportCarbonEmissionsCalculator, publicTransportCarbonEmissions.address);
	const publicTransportCarbonEmissionsCalculator = await PublicTransportCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'CarCarbonEmissions');
	const carCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(CarCarbonEmissionsCalculator, carCarbonEmissions.address);
	const carCarbonEmissionsCalculator = await CarCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'AirplaneCarbonEmissions');
	const airplaneCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(AirplaneCarbonEmissionsCalculator, airplaneCarbonEmissions.address);
	const airplaneCarbonEmissionsCalculator = await AirplaneCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'TrainCarbonEmissions');
	const trainCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(TrainCarbonEmissionsCalculator, trainCarbonEmissions.address);
	const trainCarbonEmissionsCalculator = await TrainCarbonEmissionsCalculator.deployed();

	await deployer.deploy(CarbonEmissions, 'AmericanoCarbonEmissions');
	const americanoCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(AmericanoCarbonEmissionsCalculator, americanoCarbonEmissions.address);
	const americanoCarbonEmissionsCalculator = await AmericanoCarbonEmissionsCalculator.deployed();

	console.log('BeefCarbonEmissionsCalculator : ', beefCarbonEmissionsCalculator.address);
	console.log('BicycleCarbonEmissionsCalculator : ', bicycleCarbonEmissionsCalculator.address);
	console.log('CannedTunaCarbonEmissionsCalculator : ', cannedTunaCarbonEmissionsCalculator.address);
	console.log('ChickenCarbonEmissionsCalculator : ', chickenCarbonEmissionsCalculator.address);
	console.log('EggCarbonEmissionsCalculator : ', eggCarbonEmissionsCalculator.address);
	console.log('EmailCarbonEmissionsCalculator : ', emailCarbonEmissionsCalculator.address);
	console.log('LambCarbonEmissionsCalculator : ', lambCarbonEmissionsCalculator.address);
	console.log('MilkCarbonEmissionsCalculator : ', milkCarbonEmissionsCalculator.address);
	console.log('MobileDataCarbonEmissionsCalculator : ', mobileDataCarbonEmissionsCalculator.address);
	console.log('PhoneCallCarbonEmissionsCalculator : ', phoneCallCarbonEmissionsCalculator.address);
	console.log('PorkCarbonEmissionsCalculator : ', porkCarbonEmissionsCalculator.address);
	console.log('SalmonCarbonEmissionsCalculator : ', salmonCarbonEmissionsCalculator.address);
	console.log('WalkCarbonEmissionsCalculator : ', walkCarbonEmissionsCalculator.address);
	console.log('YoutubeCarbonEmissionsCalculator : ', youtubeCarbonEmissionsCalculator.address);
	console.log('ElectricityCarbonEmissionsCalculator : ', electricityCarbonEmissionsCalculator.address);
	console.log('WaterCarbonEmissionsCalculator : ', waterCarbonEmissionsCalculator.address);
	console.log('GasCarbonEmissionsCalculator : ', gasCarbonEmissionsCalculator.address);
	console.log('PublicTransportCarbonEmissionsCalculator : ', publicTransportCarbonEmissionsCalculator.address);
	console.log('CarCarbonEmissionsCalculator : ', carCarbonEmissionsCalculator.address);
	console.log('AirplaneCarbonEmissionsCalculator : ', airplaneCarbonEmissionsCalculator.address);
	console.log('TrainCarbonEmissionsCalculator : ', trainCarbonEmissionsCalculator.address);
	console.log('AmericanoCarbonEmissionsCalculator : ', americanoCarbonEmissionsCalculator.address);
};
