// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../access/IAppIdAccessManager.sol";
import "../core/CarbonEmissions.sol";
import "./CarbonEmissionsCalculatorBase.sol";

contract CarbonHeatingCalculator is CarbonEmissionsCalculatorBase {
    //수도 탄소배출 계수 0.046946 kgCO2/MJ scaled by 10 000
    uint256 public heatingAverageCarbonEmissionValue = 469;

    address public appIdAccessManager;

    constructor(address _appIdAccessManager) {
        appIdAccessManager = _appIdAccessManager;
    }

    modifier onlyAuthorized(uint256 appId) {
        require(IAppIdAccessManager(appIdAccessManager).isAuthorized(appId, msg.sender), "AppIdAccessManager : caller is not authorized");
        _;
    }


    // calculate CO2 emission by heating in (18 decimals)
    //usedHeatingAmount scaled by 10 000
    function calculate(
        uint256 usedHeatingAmount,
        uint256 applicationId,
        address carbonEmissions,
        string memory sourceChannel
    ) external onlyAuthorized(applicationId) returns (uint256) {
        // 난방 사용량(MJ) x 난방 탄소배출 계수(kgCO2/MJ) = 탄소배출량(kgCO2)
        uint256 result = heatingAverageCarbonEmissionValue * usedHeatingAmount;

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
