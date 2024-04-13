// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";


contract ElectricityCarbonEmissionsCalculator {
    uint256 constant EMISSIONS_PER_KG = 478100000; // scaled by 10,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions)  {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value
    ) external returns (uint256) {
        uint256 result;

        result = EMISSIONS_PER_KG * (value * 10e4) / 10e4;

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}