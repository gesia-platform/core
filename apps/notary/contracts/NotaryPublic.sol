// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicAccount.sol";
import "./NotaryAccount.sol";

contract NotaryPublic is NotaryPublicAccount {
    function authorizedToApplication(
        address sender,
        uint256 applicationID
    ) public view returns (bool) {
        return
            NotaryAccount(accountsByOwner[sender]).getApplicationID() ==
            applicationID;
    }
}
