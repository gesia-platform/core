// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryAccount.sol";
import "./NotaryPublicDAO.sol";

contract NotaryPublicAccount is NotaryPublicDAO {
    uint256 private nextApplicationID = 1;

    mapping(string domain => address account) public accountsByDomain;
    mapping(address owner => address account) public accountsByOwner;
    mapping(address account => address owner) public onwersByAccount;

    mapping(address account => mapping(address notary => bool notarized))
        public accountNotarizations;

    event AccountCreated(
        address account,
        string domain,
        uint256 applicationID,
        address owner
    );

    event AccountNotarized(address account, address notary, bytes signature);

    event AccountRevoked(address account, address notary, bytes signature);

    event AccountOwnershipTransferred(
        address account,
        address previousOwner,
        address newOwner
    );

    function createAccount(
        string memory domain,
        address owner
    ) external onlyMember {
        require(
            accountsByDomain[domain] == address(0),
            "domain already in use."
        );

        NotaryAccount notaryAccount = new NotaryAccount(
            domain,
            nextApplicationID
        );

        address notaryAccountAddress = address(notaryAccount);

        accountsByDomain[domain] = notaryAccountAddress;

        accountsByOwner[owner] = notaryAccountAddress;
        onwersByAccount[notaryAccountAddress] = owner;

        emit AccountCreated(
            notaryAccountAddress,
            domain,
            nextApplicationID,
            owner
        );

        nextApplicationID++;
    }

    function notarizeAccount(
        address account,
        bytes memory signature
    ) external onlyMember {
        NotaryAccount notaryAccount = NotaryAccount(account);

        string memory domain = notaryAccount.getDomain();
        require(accountsByDomain[domain] == account, "invalid account.");

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

    function transferAccountOwnership(
        address account,
        address newOwner
    ) external onlyMember {
        bool isNotary = accountNotarizations[account][msg.sender] == true;
        bool isOwner = accountsByOwner[msg.sender] == account;

        require(isNotary || isOwner, "unauthorized.");

        address previousOwner = onwersByAccount[account];

        accountsByOwner[newOwner] = account;
        onwersByAccount[account] = newOwner;

        emit AccountOwnershipTransferred(account, previousOwner, newOwner);
    }

    function getAccountByDomain(
        string memory domain
    ) external view returns (address notaryAccount) {
        return accountsByDomain[domain];
    }

    function getAccountOwner(
        address account
    ) external view returns (address notaryAccount) {
        return onwersByAccount[account];
    }

    function getAccountNotarized(
        address account,
        address notary
    ) public view returns (bool) {
        return accountNotarizations[account][notary];
    }
}
