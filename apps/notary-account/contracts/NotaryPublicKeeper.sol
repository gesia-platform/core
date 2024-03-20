// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicDAO.sol";
import "./NotaryKeeper.sol";

contract NotaryPublicKeeper is NotaryPublicDAO {
    mapping(address keeper => mapping(address notary => bool notarized))
        public keeperNotarizations;

    mapping(address keeper => mapping(address caller => bool authorized))
        public keeperAuthorizations;

    event KeeperNotarized(address keeper, bytes signature);

    event KeeperRevoked(address keeper, bytes signature);

    event KeeperGranted(address keeper, address grantee);

    event KeeperCallerAuthorized(
        address keeper,
        address caller,
        bool authorization
    );

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

    function authorizeKeeperCaller(
        address keeper,
        address caller,
        bool authorization
    ) external {
        require(
            keeperNotarizations[keeper][msg.sender] == true,
            "sender has never notarized."
        );

        keeperAuthorizations[keeper][caller] = authorization;

        emit KeeperCallerAuthorized(keeper, caller, authorization);
    }

    function keeperCallerAuthorized(
        address keeper,
        address caller
    ) public view returns (bool) {
        return keeperAuthorizations[keeper][caller];
    }

    function getKeeperNotarized(
        address keeper,
        address notary
    ) public view returns (bool) {
        return keeperNotarizations[keeper][notary];
    }
}
