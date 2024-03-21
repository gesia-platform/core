// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract NotaryPublicDAO is Ownable {
    using ECDSA for bytes32;

    mapping(address => bool) public memberships;

    modifier onlyMember() {
        require(memberships[msg.sender] == true, "invalid dao membership.");
        _;
    }

    function setMembership(address member, bool state) external onlyOwner {
        if (state) {
            memberships[member] = true;
        } else {
            delete memberships[member];
        }
    }

    function getMembership(address member) public view returns (bool) {
        return memberships[member];
    }

    function getNotary(
        bytes memory message,
        bytes memory signature
    ) public pure returns (address) {
        return keccak256(message).toEthSignedMessageHash().recover(signature);
    }
}
