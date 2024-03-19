// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryAccount.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract NotaryPublic {
    address public owner;

    mapping(string => mapping(string => address)) public notaryAccounts; // chain > domain > account

    mapping(string => bytes32) public chainMerkleRoots; // chain > root

    event NotaryAccountCreated(
        string chain,
        string domain,
        address notaryAccount
    );

    constructor() {
        owner = msg.sender;
    }

    function notarizeRPC(
        address notaryAccount,
        bytes memory signature,
        string memory url,
        bytes32[] calldata merkleProof
    ) external {
        notaryAccount = NotaryAccount(notaryAccount);

        chain = notaryAccount.getChain();

        require(
            MerkleProof.verify(
                merkleProof,
                chainMerkleRoots[chain],
                keccak256(abi.encodePacked(msg.sender))
            ) == true,
            "invalid merkle proof within chain."
        );

        notaryAccount.notarizeRPC(signature, url);
    }

    function createNotaryAccount(
        string memory chain,
        string memory domain,
        uint verificationBlockHeight,
        bytes32[] calldata merkleProof
    ) external returns (address notaryAccount) {
        require(
            MerkleProof.verify(
                merkleProof,
                chainMerkleRoots[chain],
                keccak256(abi.encodePacked(msg.sender))
            ) == true,
            "invalid merkle proof within chain."
        );

        require(
            notaryAccounts[chain][domain] == address(0),
            "already registered domain within the chain."
        );

        notaryAccount = address(
            new NotaryAccount(
                address(this),
                verificationBlockHeight,
                chain,
                domain
            )
        );

        notaryAccounts[chain][domain] = notaryAccount;

        emit NotaryAccountCreated(chain, domain, notaryAccount);

        return notaryAccount;
    }

    function setOwner(address _owner) external {
        require(msg.sender == owner);
        owner = _owner;
    }

    function setChainMerkleRoot(
        string memory chain,
        bytes32 merkleRoot
    ) external {
        require(msg.sender == owner);
        chainMerkleRoots[chain] = merkleRoot;
    }

    function getNotaryAccount(
        string memory chain,
        string memory domain
    ) external view returns (address notaryAccount) {
        return notaryAccounts[chain][domain];
    }
}
