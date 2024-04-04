// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";

contract ElectricityCarbonEmissionsCalculator {
    uint256 constant EMISSIONS_PER_KG = 4781; // scaled by 10,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions) {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value,
        bool isSummer
    ) external returns (uint256) {
        uint256 result;

        if (isSummer) {
            result = EMISSIONS_PER_KG * (value * 10e4) * 2 / 10e4;
        } else {
            result = EMISSIONS_PER_KG * (value * 10e4);
        }

        carbonEmissions.mint(msg.sender, applicationID, result, userID);

        return result;
    }
}