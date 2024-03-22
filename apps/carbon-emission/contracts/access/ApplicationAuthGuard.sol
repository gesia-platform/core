// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./IApplicationAccessManager.sol";

abstract contract ApplicationAuthGuard {
    IApplicationAccessManager public immutable applicationAccessManager;

    constructor(address _applicationAccessManager) {
        applicationAccessManager = IApplicationAccessManager(
            _applicationAccessManager
        );
    }

    error ApplicationUnauthorized(uint256 applicationID, address sender);

    modifier onlyAuthorized(uint256 applicationID) {
        if (!applicationAccessManager.isAuthorized(applicationID, msg.sender)) {
            revert ApplicationUnauthorized(applicationID, msg.sender);
        }
        _;
    }
}
