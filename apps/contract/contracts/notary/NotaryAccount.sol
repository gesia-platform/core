// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

import "./NotaryPublic.sol";

contract NotaryAccount {
    NotaryPublic public immutable notaryPublic;

    constructor(NotaryPublic _notaryPublic) {
        notaryPublic = _notaryPublic;
    }

    function notaryCall(
        address target,
        bytes memory data
    ) external returns (bool success, bytes memory result) {
        require(
            notaryPublic.hasMembership(msg.sender),
            "sender is not notary public member"
        );

        return target.call(data);
    }
}
