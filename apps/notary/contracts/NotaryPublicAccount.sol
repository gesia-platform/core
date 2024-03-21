// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryAccount.sol";
import "./NotaryPublicDAO.sol";

contract NotaryPublicAccount is NotaryPublicDAO {
    uint256 private nextApplicationID = 1;

    mapping(string domain => address account) public accounts;

    mapping(address account => mapping(address notary => bool notarized))
        public accountNotarizations;

    event AccountCreated(address account, string domain, uint256 applicationID);

    event AccountNotarized(address account, address notary, bytes signature);

    event AccountRevoked(address account, address notary, bytes signature);

    event AccountOwnerUpdated(
        address account,
        address notary,
        address newOwner
    );

    function createAccount(
        string memory domain,
        address owner
    ) external onlyMember {
        require(accounts[domain] == address(0), "domain already in use.");

        NotaryAccount notaryAccount = new NotaryAccount(
            address(this),
            domain,
            nextApplicationID,
            owner
        );

        accounts[domain] = address(notaryAccount);

        emit AccountCreated(address(notaryAccount), domain, nextApplicationID);

        nextApplicationID++;
    }

    function notarizeAccount(
        address account,
        bytes memory signature
    ) external onlyMember {
        NotaryAccount notaryAccount = NotaryAccount(account);

        string memory domain = notaryAccount.getDomain();
        require(accounts[domain] == account, "invalid account.");

        bytes memory message = abi.encodePacked("notarize_account", account);
        address notary = getNotary(message, signature);
        require(notary == msg.sender, "sender is not the notary.");

        accountNotarizations[account][notary] = true;

        emit AccountNotarized(address(notaryAccount), notary, signature);
    }

    function revokeAccount(address account, bytes memory signature) external {
        require(
            accountNotarizations[account][msg.sender] == true,
            "sender is not notary."
        );

        NotaryAccount notaryAccount = NotaryAccount(account);

        bytes memory message = abi.encodePacked("revoke_account", account);
        address notary = getNotary(message, signature);
        require(notary == msg.sender, "sender is not the notary.");

        delete accountNotarizations[account][notary];

        emit AccountRevoked(address(notaryAccount), notary, signature);
    }

    function setAccountOwner(
        address account,
        address newOwner
    ) external onlyMember {
        require(
            accountNotarizations[account][msg.sender] == true,
            "sender is not notary."
        );

        NotaryAccount notaryAccount = NotaryAccount(account);

        notaryAccount.setOwner(newOwner);

        emit AccountOwnerUpdated(address(notaryAccount), msg.sender, newOwner);
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