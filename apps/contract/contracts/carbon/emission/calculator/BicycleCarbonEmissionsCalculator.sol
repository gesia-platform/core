// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "../CarbonEmissions.sol";
import "../../../notary/NotaryModule.sol";
import "../util/SafeMath.sol";

contract BicycleCarbonEmissionsCalculator is NotaryModule {
    using SafeMath for uint256;

    uint256 EMISSIONS_PER_HOUR = 46900000; // scaled by 1,000,000,000

    CarbonEmissions public immutable carbonEmissions;

    constructor(CarbonEmissions _carbonEmissions, NotaryPublic _notaryPublic) NotaryModule(_notaryPublic) {
        carbonEmissions = _carbonEmissions;
    }

    function calculate(
        uint256 applicationID,
        bytes memory userID,
        uint256 value
    ) external returns (uint256) {
        uint256 result = EMISSIONS_PER_HOUR.mul(value).mul(10e4).div(10e4);

        carbonEmissions.mint(applicationID, result, userID);

        return result;
    }
}
