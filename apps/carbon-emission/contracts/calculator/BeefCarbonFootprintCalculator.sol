// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.23;

import "../access/ApplicationAuthGuard.sol";
import "../CarbonEmissions.sol";

contract BeefCarbonFootprintCalculator is ApplicationAuthGuard {
    uint256 EMISSIONS_PER_KG = 270000; // scaled by 10,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(
        address _applicationAccessManager,
        CarbonEmissions _carbonEmissions
    ) ApplicationAuthGuard(_applicationAccessManager) {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value
    ) external onlyAuthorized(applicationID) returns (uint256) {
        uint256 result = EMISSIONS_PER_KG * (value * 10e4);

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}
