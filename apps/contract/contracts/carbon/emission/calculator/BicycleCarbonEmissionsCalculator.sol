// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";
import "../../../notary/NotaryModule.sol";

contract BicycleCarbonEmissionsCalculator is NotaryModule {
    uint256 EMISSIONS_PER_HOUR = 2813900000; // scaled by 1,000,000,000
    uint256 constant MINUTE_TO_HOUR = 60;

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions, NotaryPublic _notaryPublic) NotaryModule(_notaryPublic) {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value
    ) external returns (uint256) {
        uint256 result = EMISSIONS_PER_HOUR * (value * 10e4) / 10e4 / MINUTE_TO_HOUR;

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}
