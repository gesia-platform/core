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
const PaperCarbonEmissionsCalculator = artifacts.require('PaperCarbonEmissionsCalculator');

module.exports = async (deployer) => {
	if (process.env.SKIP_MIGRATIONS) return;

	await deployer.deploy(CarbonEmissions, 'BeefCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const beefCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(BeefCarbonEmissionsCalculator, beefCarbonEmissions.address);
	const beefCarbonEmissionsCalculator = await BeefCarbonEmissionsCalculator.deployed();
	await beefCarbonEmissions.setCalculatorApproval(beefCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'BicycleCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const bicycleCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(BicycleCarbonEmissionsCalculator, bicycleCarbonEmissions.address);
	const bicycleCarbonEmissionsCalculator = await BicycleCarbonEmissionsCalculator.deployed();
	await bicycleCarbonEmissions.setCalculatorApproval(bicycleCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'CannedTunaCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const cannedTunaCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(CannedTunaCarbonEmissionsCalculator, cannedTunaCarbonEmissions.address);
	const cannedTunaCarbonEmissionsCalculator = await CannedTunaCarbonEmissionsCalculator.deployed();
	await cannedTunaCarbonEmissions.setCalculatorApproval(cannedTunaCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'CannedTunaCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const chickenCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(ChickenCarbonEmissionsCalculator, chickenCarbonEmissions.address);
	const chickenCarbonEmissionsCalculator = await ChickenCarbonEmissionsCalculator.deployed();
	await chickenCarbonEmissions.setCalculatorApproval(chickenCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'EggCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const eggCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(EggCarbonEmissionsCalculator, eggCarbonEmissions.address);
	const eggCarbonEmissionsCalculator = await EggCarbonEmissionsCalculator.deployed();
	await eggCarbonEmissions.setCalculatorApproval(eggCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'EmailCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const emailCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(EmailCarbonEmissionsCalculator, emailCarbonEmissions.address);
	const emailCarbonEmissionsCalculator = await EmailCarbonEmissionsCalculator.deployed();
	await emailCarbonEmissions.setCalculatorApproval(emailCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'LambCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const lambCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(LambCarbonEmissionsCalculator, lambCarbonEmissions.address);
	const lambCarbonEmissionsCalculator = await LambCarbonEmissionsCalculator.deployed();
	await lambCarbonEmissions.setCalculatorApproval(lambCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'MilkCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const milkCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(MilkCarbonEmissionsCalculator, milkCarbonEmissions.address);
	const milkCarbonEmissionsCalculator = await MilkCarbonEmissionsCalculator.deployed();
	await milkCarbonEmissions.setCalculatorApproval(milkCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'MobileCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const mobileCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(MobileDataCarbonEmissionsCalculator, mobileCarbonEmissions.address);
	const mobileDataCarbonEmissionsCalculator = await MobileDataCarbonEmissionsCalculator.deployed();
	await mobileCarbonEmissions.setCalculatorApproval(mobileDataCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'PhoneCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const phoneCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(PhoneCallCarbonEmissionsCalculator, phoneCarbonEmissions.address);
	const phoneCallCarbonEmissionsCalculator = await PhoneCallCarbonEmissionsCalculator.deployed();
	await phoneCarbonEmissions.setCalculatorApproval(phoneCallCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'PorkCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const porkCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(PorkCarbonEmissionsCalculator, porkCarbonEmissions.address);
	const porkCarbonEmissionsCalculator = await PorkCarbonEmissionsCalculator.deployed();
	await porkCarbonEmissions.setCalculatorApproval(porkCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'SalmonCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const salmonCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(SalmonCarbonEmissionsCalculator, salmonCarbonEmissions.address);
	const salmonCarbonEmissionsCalculator = await SalmonCarbonEmissionsCalculator.deployed();
	await salmonCarbonEmissions.setCalculatorApproval(salmonCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'WalkCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const walkCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(WalkCarbonEmissionsCalculator, walkCarbonEmissions.address);
	const walkCarbonEmissionsCalculator = await WalkCarbonEmissionsCalculator.deployed();
	await walkCarbonEmissions.setCalculatorApproval(walkCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'YoutubeCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const youtubeCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(YoutubeCarbonEmissionsCalculator, youtubeCarbonEmissions.address);
	const youtubeCarbonEmissionsCalculator = await YoutubeCarbonEmissionsCalculator.deployed();
	await youtubeCarbonEmissions.setCalculatorApproval(youtubeCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'ElectricityCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const electricityCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(ElectricityCarbonEmissionsCalculator, electricityCarbonEmissions.address);
	const electricityCarbonEmissionsCalculator = await ElectricityCarbonEmissionsCalculator.deployed();
	await electricityCarbonEmissions.setCalculatorApproval(electricityCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'WaterCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const waterCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(WaterCarbonEmissionsCalculator, waterCarbonEmissions.address);
	const waterCarbonEmissionsCalculator = await WaterCarbonEmissionsCalculator.deployed();
	await waterCarbonEmissions.setCalculatorApproval(waterCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'GasCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const gasCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(GasCarbonEmissionsCalculator, gasCarbonEmissions.address);
	const gasCarbonEmissionsCalculator = await GasCarbonEmissionsCalculator.deployed();
	await gasCarbonEmissions.setCalculatorApproval(gasCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'PublicTransportCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const publicTransportCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(PublicTransportCarbonEmissionsCalculator, publicTransportCarbonEmissions.address);
	const publicTransportCarbonEmissionsCalculator = await PublicTransportCarbonEmissionsCalculator.deployed();
	await publicTransportCarbonEmissions.setCalculatorApproval(publicTransportCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'CarCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const carCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(CarCarbonEmissionsCalculator, carCarbonEmissions.address);
	const carCarbonEmissionsCalculator = await CarCarbonEmissionsCalculator.deployed();
	await carCarbonEmissions.setCalculatorApproval(carCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'AirplaneCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const airplaneCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(AirplaneCarbonEmissionsCalculator, airplaneCarbonEmissions.address);
	const airplaneCarbonEmissionsCalculator = await AirplaneCarbonEmissionsCalculator.deployed();
	await airplaneCarbonEmissions.setCalculatorApproval(airplaneCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'TrainCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const trainCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(TrainCarbonEmissionsCalculator, trainCarbonEmissions.address);
	const trainCarbonEmissionsCalculator = await TrainCarbonEmissionsCalculator.deployed();
	await trainCarbonEmissions.setCalculatorApproval(trainCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'AmericanoCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const americanoCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(AmericanoCarbonEmissionsCalculator, americanoCarbonEmissions.address);
	const americanoCarbonEmissionsCalculator = await AmericanoCarbonEmissionsCalculator.deployed();
	await americanoCarbonEmissions.setCalculatorApproval(americanoCarbonEmissionsCalculator.address, true);

	await deployer.deploy(CarbonEmissions, 'PaperCarbonEmissions', process.env.EMISSION_NOTARY_PUBLIC);
	const paperCarbonEmissions = await CarbonEmissions.deployed();
	await deployer.deploy(PaperCarbonEmissionsCalculator, paperCarbonEmissions.address);
	const paperCarbonEmissionsCalculator = await PaperCarbonEmissionsCalculator.deployed();
	await paperCarbonEmissions.setCalculatorApproval(paperCarbonEmissionsCalculator.address, true);

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
	console.log('PaperCarbonEmissionsCalculator : ', paperCarbonEmissionsCalculator.address);
};
