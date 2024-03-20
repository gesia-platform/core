// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryAccount.sol";
import "./NotaryPublicDAO.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

contract NotaryPublicAccount is NotaryPublicDAO {
    mapping(string domain => address account) public accounts;

    mapping(address account => mapping(address notary => bool notarized))
        public accountNotarizations;

    event AccountCreated(address account, string domain);

    event AccountOwnerUpdated(address account, address newOwner);

    event AccountNotarized(address account, bytes signature);

    event AccountCanceled(address account, bytes signature);

    /** Manage Accounts */

    function notarizeAccount(
        address account,
        bytes memory signature,
        string memory url,
        bytes32[] calldata merkleProof
    ) external onlyMerkleProof(merkleProof) {
        NotaryAccount notaryAccount = NotaryAccount(account);

        string memory domain = notaryAccount.getDomain();

        require(
            accounts[domain] == account,
            "this account not managed by notary public."
        );

        address notary = getNotary(abi.encodePacked(domain), signature);

        require(notary == msg.sender, "notraize must it self.");

        accountNotarizations[account][notary] = true;

        notaryAccount.notarize(notary, url);

        emit AccountNotarized(address(notaryAccount), signature);
    }

    function cancelAccount(address account, bytes memory signature) external {
        require(
            accountNotarizations[account][msg.sender] == true,
            "sender has never notarized."
        );

        NotaryAccount notaryAccount = NotaryAccount(account);

        string memory domain = notaryAccount.getDomain();

        address notary = getNotary(abi.encodePacked(domain), signature);

        require(notary == msg.sender, "notraize must it self.");

        delete accountNotarizations[account][notary];

        notaryAccount.cancel(notary);

        emit AccountCanceled(address(notaryAccount), signature);
    }

    function createAccount(
        string memory domain,
        address _owner,
        bytes32[] calldata merkleProof
    ) external onlyMerkleProof(merkleProof) {
        require(accounts[domain] == address(0), "already created domain.");

        address account = address(
            new NotaryAccount(address(this), domain, _owner)
        );

        accounts[domain] = account;

        emit AccountCreated(account, domain);
    }

    function setAccountOwner(
        address account,
        address newOwner,
        bytes32[] calldata merkleProof
    ) external onlyMerkleProof(merkleProof) {
        require(
            accountNotarizations[account][msg.sender] == true,
            "sender has never notarized."
        );

        NotaryAccount notaryAccount = NotaryAccount(account);

        notaryAccount.setOwner(newOwner);

        emit AccountOwnerUpdated(address(notaryAccount), newOwner);
    }

    function getAccount(
        string memory domain
    ) external view returns (address notaryAccount) {
        return accounts[domain];
    }

    function getAccountNotarized(
        address account,
        address notary
    ) public view returns (bool) {
        return accountNotarizations[account][notary];
    }
}
