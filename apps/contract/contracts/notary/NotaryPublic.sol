// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

contract NotaryPublic {
    mapping(bytes1 prefix => mapping(address notary => bytes signature))
        internal notarizations;

    event Notarized(bytes1 prefix, bytes signatrue, address notary);

    function notarize(bytes1 prefix, bytes memory signatrue) external {
        notarizations[prefix][msg.sender] = signatrue;

        emit Notarized(prefix, signatrue, msg.sender);
    }
}
