// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";


contract TrainCarbonEmissionsCalculator {
    uint256 constant COMMON_EMISSIONS_PER_KG = 22560290; // scaled by 1,000,000,000
    uint256 constant KTX_SRT_EMISSIONS_PER_KG = 25334000; // scaled by 1,000,000,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions)  {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value,
        bool isCommon
    ) external returns (uint256) {
        uint256 result;

        if(isCommon == true) {
            result =  COMMON_EMISSIONS_PER_KG * (value * 10e4) / 10e4;
        } else {
            result =  KTX_SRT_EMISSIONS_PER_KG * (value * 10e4) / 10e4;
        }

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}