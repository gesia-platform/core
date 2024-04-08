// SPDX-License-Identifier: MIT 
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";
import "../../../notary/NotaryModule.sol";

contract PublicTransportCarbonEmissionsCalculator is NotaryModule {
    uint256 constant BUS_EMISSIONS_PER_KG = 5743043; // scaled by 10,000,000,000
    uint256 constant SUBWAY_EMISSIONS_PER_KG = 43729; // scaled by 10,000,000,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions, NotaryPublic _notaryPublic) NotaryModule(_notaryPublic) {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value,
        bool isBus
    ) external returns (uint256) {
        uint256 result;

        if (isBus) {
            result = BUS_EMISSIONS_PER_KG * (value * 10e4) / 10e4;
        } else {
            result = SUBWAY_EMISSIONS_PER_KG * (value * 10e4) / 10e4;
        }

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}