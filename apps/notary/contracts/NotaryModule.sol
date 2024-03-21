// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublic.sol";

abstract contract NotaryModule {
    NotaryPublic private immutable notaryPublic;

    constructor(NotaryPublic _notaryPublic) {
        notaryPublic = _notaryPublic;
    }

    modifier onlyAuthorized() {
        require(
            notaryPublic.moduleCallAuthorized(address(this), msg.sender),
            "sender is not authorized to module."
        );
        _;
    }

    function notaryPublicAddress() public view returns (address) {
        return address(notaryPublic);
    }
}
