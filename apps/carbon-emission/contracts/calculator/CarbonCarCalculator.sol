// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../operator/IOperator.sol";
import "../core/CarbonEmissions.sol";
import "./CarbonEmissionsCalculatorBase.sol";

contract CarbonCarCalculator is CarbonEmissionsCalculatorBase {
    mapping(string => uint256) private fuelEmissions;

    address public operatorManager;

    modifier operatorsOnly() {
        require(IOperator(operatorManager).isOperator(msg.sender), "#operatorsOnly:");
        _;
    }

    constructor(address _operatorManager){
        operatorManager = _operatorManager;
    }

    function addFuelVolumes(string memory fuelType, uint256 emission) external operatorsOnly {
        //fuel emission (탄소배출계수(kgCO2/L)) scaled by 10 000
        fuelEmissions[fuelType] = emission;
    }

    // get the emission for a given fuel type
    function getFuelEmission(string memory fuelType) public view returns (uint256) {
        require(fuelEmissions[fuelType] > 0, "Invalid fuel type");
        return fuelEmissions[fuelType];
    }

    // calculate CO2 emission by fuel type in (18 decimals)
    function calculateCO2ByFuel(
        string memory fuelType, 
        uint256 distance, 
        uint256 efficiency,
        uint256 applicationId,
        address carbonEmissions,
        string memory sourceChannel
    ) external returns (uint256) {
        //주행 거리(km) / 연비(km/L) x 연료별 탄소 배출 계수(kgCO2/L) = 배출량(kgCO2)
        uint256 emission = getFuelEmission(fuelType);
        uint256 fuelUsed = (distance / efficiency);

        uint256 result = (fuelUsed * emission * 1e18) / 1e4;

        onCalculated(
            carbonEmissions,
            msg.sender,
            result,
            applicationId,
            sourceChannel
        );

        return result;
    }

    function onCalculated(
        address carbonEmissions,
        address from,
        uint256 amount,
        uint256 applicationId,
        string memory sourceChannel
    ) internal virtual override {
        CarbonEmissions(carbonEmissions).mint(
            from,
            amount,
            applicationId,
            sourceChannel
        );
    }
}
