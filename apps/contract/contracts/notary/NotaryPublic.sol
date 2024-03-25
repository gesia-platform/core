// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

contract NotaryPublic {
    mapping(bytes1 prefix => mapping(address notary => bytes signature))
        internal notarizations;

    mapping(address notary => bytes pubkey) internal pubkeys;

    event Notarized(bytes1 prefix, bytes signatrue, address notary);

    function notarize(bytes1 prefix, bytes memory signatrue) external {
        notarizations[prefix][msg.sender] = signatrue;

        emit Notarized(prefix, signatrue, msg.sender);
    }

    function register(bytes memory pubkey) external {
        pubkeys[msg.sender] = pubkey;
    }

    function getPubkey(address notary) public view returns (bytes memory) {
        return pubkeys[notary];
    }

    function getNotarization(
        bytes1 prefix,
        address notary
    ) public view returns (bytes memory, bytes memory) {
        return (pubkeys[notary], notarizations[prefix][notary]);
    }
}
