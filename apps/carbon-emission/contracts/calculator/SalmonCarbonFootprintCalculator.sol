// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.23;


import "../CarbonEmissions.sol";

contract BicycleCarbonEmissionsCalculator {
    uint256 EMISSIONS_PER_KG = 119000; // scaled by 10,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(
        CarbonEmissions _carbonEmissions
    ) {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value
    ) external returns (uint256) {
        uint256 result = EMISSIONS_PER_KG * (value * 10e4);

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}
