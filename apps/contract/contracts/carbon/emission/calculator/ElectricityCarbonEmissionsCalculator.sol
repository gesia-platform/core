// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";

contract ElectricityCarbonEmissionsCalculator {
    uint256 constant EMISSIONS_PER_KG = 47810; // scaled by 100,000 = 0.4781

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
            result = EMISSIONS_PER_KG * (value * 10e5) * 2;
        } else {
            result = EMISSIONS_PER_KG * (value * 10e5);
        }

        carbonEmissions.mint(msg.sender, applicationID, result, userID);

        return result;
    }
}