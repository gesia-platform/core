// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract NotaryPublicDAO is Ownable {
    using ECDSA for bytes32;

    bytes32 public merkleRoot;

    modifier onlyMerkleProof(bytes32[] calldata merkleProof) {
        require(
            MerkleProof.verify(
                merkleProof,
                merkleRoot,
                keccak256(abi.encodePacked(msg.sender))
            ) == true,
            "invalid merkle proof."
        );
        _;
    }

    function setMerkleRoot(bytes32 _merkleRoot) external onlyOwner {
        merkleRoot = _merkleRoot;
    }

    function getMerkleRoot() external view returns (bytes32) {
        return merkleRoot;
    }

    function getNotary(
        bytes memory message,
        bytes memory signature
    ) public pure returns (address) {
        return keccak256(message).toEthSignedMessageHash().recover(signature);
    }
}
