// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";

contract EggCarbonEmissionsCalculator {
    uint256 EMISSIONS_PER_DOZEN = 240; // scaled by 10,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions) {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value
    ) external returns (uint256) {
        uint256 result = EMISSIONS_PER_DOZEN * (value * 10e4);

        carbonEmissions.mint(msg.sender, applicationID, result, userID);

        return result;
    }
}
