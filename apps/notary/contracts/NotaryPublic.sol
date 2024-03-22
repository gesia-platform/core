// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract NotaryPublic is Ownable {
    using ECDSA for bytes32;

    struct Application {
        uint256 id;
        string domain;
        address account;
    }

    struct ApplicationNotarization {
        bool notarized;
        bytes signature;
        string url;
    }

    uint256 private nextApplicationID = 1;

    mapping(uint256 => Application) internal applications;

    mapping(uint256 => mapping(string => mapping(address => ApplicationNotarization)))
        internal applicationNotarizations;

    event ApplicationCreated(
        uint256 applicationID,
        string domain,
        address account
    );

    event ApplicationNotarized(
        uint256 applicationID,
        string network,
        address notary
    );

    event ApplicationNotarizationRevoked(
        uint256 applicationID,
        string network,
        address notary
    );

    function createApplication(string memory domain, address account) external {
        applications[nextApplicationID] = Application({
            id: nextApplicationID,
            domain: domain,
            account: account
        });

        emit ApplicationCreated(nextApplicationID, domain, account);
        nextApplicationID++;
    }

    function notarizeApplication(
        uint256 applicationID,
        string memory network,
        string memory url,
        bytes memory signature
    ) external {
        require(applications[applicationID].id == 0, "invalid application.");

        bytes memory message = abi.encodePacked(
            "notarize_application",
            applicationID
        );

        address notary = _getNotary(message, signature);

        require(notary == msg.sender, "sender is not the notary.");

        applicationNotarizations[applicationID][network][
            notary
        ] = ApplicationNotarization(true, signature, url);

        emit ApplicationNotarized(applicationID, network, notary);
    }

    function revokeApplicationNotarization(
        uint256 applicationID,
        string memory network,
        bytes memory signature
    ) external {
        Application memory application = applications[applicationID];

        require(application.id == 0, "invalid application.");

        bytes memory message = abi.encodePacked(
            "revoke_application_notarization",
            applicationID
        );

        address notary = _getNotary(message, signature);

        ApplicationNotarization storage notarization = applicationNotarizations[
            applicationID
        ][network][notary];

        require(notarization.notarized == true, "notary did not notarize.");

        notarization.notarized = false;

        notarization.signature = signature;

        notarization.url = "";

        emit ApplicationNotarizationRevoked(applicationID, network, notary);
    }

    function getApplicationDetails(
        uint256 applicationID
    ) public view returns (string memory, address) {
        Application memory application = applications[applicationID];

        return (application.domain, application.account);
    }

    function getApplicationNotarizationDetails(
        uint256 applicationID,
        string memory network,
        address notary
    ) public view returns (bool, bytes memory, string memory) {
        ApplicationNotarization memory notarization = applicationNotarizations[
            applicationID
        ][network][notary];

        return (
            notarization.notarized,
            notarization.signature,
            notarization.url
        );
    }

    function _getNotary(
        bytes memory message,
        bytes memory signature
    ) internal pure returns (address) {
        return keccak256(message).toEthSignedMessageHash().recover(signature);
    }
}
