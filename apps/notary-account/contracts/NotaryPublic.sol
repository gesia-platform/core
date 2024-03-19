// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryAccount.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract NotaryPublic {
    address public owner;

    mapping(string => address) public notaryAccounts;

    bytes32 public merkleRoot;

    event NotaryAccountUpdated(
        string domain,
        address caller,
        address notaryAccount
    );

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

    constructor() {
        owner = msg.sender;
    }

    function notarizeAccount(
        address notaryAccount,
        bytes memory signature,
        string memory url,
        bytes32[] calldata merkleProof
    ) external onlyMerkleProof(merkleProof) {
        NotaryAccount(notaryAccount).notarize(signature, url);
    }

    function createAccount(
        string memory domain,
        address caller,
        bytes32[] calldata merkleProof
    ) external onlyMerkleProof(merkleProof) {
        require(
            notaryAccounts[domain] == address(0),
            "already created domain."
        );

        address notaryAccount = address(
            new NotaryAccount(address(this), domain, caller)
        );

        notaryAccounts[domain] = notaryAccount;

        emit NotaryAccountUpdated(domain, caller, notaryAccount);
    }

    function setAccountCaller(
        address notaryAccount,
        address caller,
        bytes32[] calldata merkleProof
    ) external onlyMerkleProof(merkleProof) {
        NotaryAccount(notaryAccount).setCaller(caller);
    }

    function setOwner(address _owner) external {
        require(msg.sender == owner);
        owner = _owner;
    }

    function setMerkleRoot(bytes32 _merkleRoot) external {
        require(msg.sender == owner);
        merkleRoot = _merkleRoot;
    }

    function getMerkleRoot() external view returns (bytes32) {
        return merkleRoot;
    }

    function getOwner() external view returns (address) {
        return owner;
    }

    function getNotaryAccount(
        string memory domain
    ) external view returns (address notaryAccount) {
        return notaryAccounts[domain];
    }
}
