// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicDAO.sol";
import "./NotaryKeeper.sol";
import "./NotaryAccount.sol";

contract NotaryPublicKeeper is NotaryPublicDAO {
    mapping(address keeper => address account) public keeperRequests;

    mapping(address keeper => mapping(address notary => bool notarized))
        public keeperNotarizations;

    mapping(address keeper => mapping(address account => bool authorized))
        public keeperAuthorizations;

    event KeeperNotarizationRequested(address keeper, address account);

    event KeeperNotarized(address keeper, bytes signature);

    event KeeperRevoked(address keeper, bytes signature);

    function requestKeeperNotarization(
        address keeper,
        address account
    ) external {
        require(
            keeperRequests[keeper] == address(0),
            "already requested keeper."
        );

        require(
            NotaryAccount(account).getOwner() == msg.sender,
            "sender is not the notary account owner."
        );

        keeperRequests[keeper] = account;

        emit KeeperNotarizationRequested(keeper, account);
    }

    function notarizeKeeper(
        address keeper,
        bytes memory signature
    ) external onlyMember {
        require(
            NotaryKeeper(keeper).getNotaryPublic() == address(this),
            "keeper granter must be notary public."
        );

        bytes memory message = abi.encodePacked("notarize_keeper", keeper);

        address notary = getNotary(message, signature);

        require(notary == msg.sender, "notraize must it self.");

        keeperNotarizations[keeper][notary] = true;

        emit KeeperNotarized(keeper, signature);
    }

    function revokeKeeper(address keeper, bytes memory signature) external {
        require(
            keeperNotarizations[keeper][msg.sender] == true,
            "sender has never notarized."
        );

        bytes memory message = abi.encodePacked("revoke_keeper", keeper);

        address notary = getNotary(message, signature);

        require(notary == msg.sender, "notraize must it self.");

        delete keeperNotarizations[keeper][notary];

        emit KeeperRevoked(keeper, signature);
    }

    function authorizeKeeper(
        address keeper,
        address account,
        bool authorization
    ) external {
        require(
            keeperNotarizations[keeper][msg.sender] == true,
            "sender has never notarized."
        );

        keeperAuthorizations[keeper][account] = authorization;
    }

    function keeperCallAuthorized(
        address keeper,
        address account,
        address caller
    ) public view returns (bool) {
        return
            keeperAuthorizations[keeper][account] &&
            NotaryAccount(account).getOwner() == caller;
    }

    function getKeeperNotarized(
        address keeper,
        address notary
    ) public view returns (bool) {
        return keeperNotarizations[keeper][notary];
    }

    function getKeeperAuthorized(
        address keeper,
        address account
    ) public view returns (bool) {
        return keeperAuthorizations[keeper][account];
    }
}
