// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../access/IAppIdAccessManager.sol";
import "../core/CarbonEmissions.sol";
import "./CarbonEmissionsCalculatorBase.sol";

contract CarbonBeefCalculator is CarbonEmissionsCalculatorBase {
    uint256 public importedCo2Emission = 30200; // scaled by 10000 // 국내산 소고기 배출계수 (1t/yr x 53GJ/t x 0.0537tCO2/GJ = 약 2.85tCO2/yr -> 약 2.85kgCO2/kg)
    uint256 public domesticCo2Emission = 28500; // scaled by 10000 // 수입 소고기 배출계수 (1t/yr x 53GJ/t x 0.0566tCO2/GJ = 약 3.02tCO2/yr -> 약 3.02kgCO2/kg)

    address public appIdAccessManager;

    constructor(address _appIdAccessManager) {
        appIdAccessManager = _appIdAccessManager;
    }

    modifier onlyAuthorized(uint256 appId) {
        require(IAppIdAccessManager(appIdAccessManager).isAuthorized(appId, msg.sender), "AppIdAccessManager : caller is not authorized");
        _;
    }


    // calculate CO2 emission by imported beef in (18 decimals)
    function calculateCO2ByImportedBeef(
            uint256 consumption,
            uint256 applicationId, 
            address carbonEmissions, 
            string memory sourceChannel
        ) external onlyAuthorized(applicationId) returns (uint256) {
        // 연간 소비 소고기(kg/yr)
        // consumption scaled up by 10000
        uint256 result = ((consumption * importedCo2Emission * 10e18)) / 10e8;

        onCalculated(
            carbonEmissions,
            msg.sender,
            result,
            applicationId,
            sourceChannel
        );
        
        return result;
    }

    // calculate CO2 emission by domestic beef in (18 decimals)
    function calculateCO2ByDomesticBeef(
            uint256 consumption, 
            uint256 applicationId, 
            address carbonEmissions, 
            string memory sourceChannel
        ) external onlyAuthorized(applicationId) returns (uint256) {
        // 연간 소비한 국내산 소고기(kg/yr)
        // weight scaled up by 10000
        uint256 result = ((consumption * domesticCo2Emission * 10e18)) / 10e8;

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
