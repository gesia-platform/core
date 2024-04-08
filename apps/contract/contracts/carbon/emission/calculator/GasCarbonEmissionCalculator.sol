// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";
import "../../../notary/NotaryModule.sol";

contract GasCarbonEmissionsCalculator is NotaryModule {
    uint256 constant EMISSIONS_PER_KG = 21760; // scaled by 10,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions, NotaryPublic _notaryPublic) NotaryModule(_notaryPublic) {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value
    ) external returns (uint256) {
        uint256 result = EMISSIONS_PER_KG * (value * 10e4) / 10e4;

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}