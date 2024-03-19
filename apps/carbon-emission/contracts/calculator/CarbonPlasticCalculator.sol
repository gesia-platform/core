// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../core/CarbonEmissions.sol";
import "./CarbonEmissionsCalculatorBase.sol";

contract CarbonPlasticCalculator is CarbonEmissionsCalculatorBase {
    uint256 public plasticCo2Emission = 68650; // PET 병 1kg 탄소배출계수(kgCO2/kg)

    // calculate CO2 emission by plastic in (18 decimals)
    function calculateCO2ByPlastic(
        uint256 weight,
        uint256 applicationId,
        address carbonEmissions,
        string memory sourceChannel
    ) external returns (uint256) {
        // 플라스틱(PET) 무게(kg) x PET 병 1kg 탄소배출계수(kgCO2/kg) = 배출량(kgCO2)
        // weight scaled up by 10000
        uint256 result =  ((weight * plasticCo2Emission) / 1e18) * 1e8;

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