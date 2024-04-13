// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";


contract AirplaneCarbonEmissionsCalculator {
    uint256 constant EMISSIONS_PER_KG = 3160000000; // scaled by 1,000,000,000
    uint256 constant AVERAGE_VALUE = 16; // scaled by 1,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions)  {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value
    ) external returns (uint256) {
        uint256 result = EMISSIONS_PER_KG * AVERAGE_VALUE * (value * 10e4) / 10e4 / 10e2;

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}