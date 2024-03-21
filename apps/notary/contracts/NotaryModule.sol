// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublic.sol";
import "./NotaryKeeper.sol";

abstract contract NotaryModule {
    NotaryPublic private immutable notaryPublic;

    constructor(NotaryPublic _notaryPublic) {
        notaryPublic = _notaryPublic;
    }

    modifier onlyAuthorized() {
        require(
            notaryPublic.authorizedToModule(msg.sender, address(this)) == true,
            "sender is not authorized to module."
        );
        _;
    }

    modifier onlyAuthorizedApplication(uint256 applicationID) {
        require(
            notaryPublic.authorizedToApplication(msg.sender, applicationID) ==
                true,
            "sender is not authorized to application."
        );
        _;
    }

    function notaryPublicAddress() public view returns (address) {
        return address(notaryPublic);
    }
}
