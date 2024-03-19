// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../access/IAppIdAccessManager.sol";
import "../core/CarbonEmissions.sol";
import "./CarbonEmissionsCalculatorBase.sol";

contract CarbonWaterCalculator is CarbonEmissionsCalculatorBase {
    //수도 탄소배출 계수 0.247kgCO2/m³ scaled by 10 000
    uint256 public waterCarbonEmissionValue = 2470;

    address public appIdAccessManager;

    constructor(address _appIdAccessManager) {
        appIdAccessManager = _appIdAccessManager;
    }

    modifier onlyAuthorized(uint256 appId) {
        require(IAppIdAccessManager(appIdAccessManager).isAuthorized(appId, msg.sender), "AppIdAccessManager : caller is not authorized");
        _;
    }


    // calculate CO2 emission by water in (18 decimals)
    function calculateCO2ByWater(
        uint256 usedWaterAmount,
        uint256 applicationId,
        address carbonEmissions,
        string memory sourceChannel
    ) external onlyAuthorized(applicationId) returns (uint256) {
        // 수도 사용량(m³) x 수도 탄소배출 계수(kgCO2/m³) = 탄소배출량(kgCO2)
        uint256 result = waterCarbonEmissionValue * usedWaterAmount;

        uint256 scaledResult = (result * 1e18) / 1e8;

        onCalculated(
            carbonEmissions,
            msg.sender,
            scaledResult,
            applicationId,
            sourceChannel
        );

        return scaledResult;
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
