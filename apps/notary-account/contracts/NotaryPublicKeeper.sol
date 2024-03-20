// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicDAO.sol";
import "./NotaryKeeper.sol";

contract NotaryPublicKeeper is NotaryPublicDAO {
    mapping(address keeper => mapping(address notary => bool notarized))
        public keeperNotarizations;

    event KeeperNotarized(address keeper, bytes signature);

    event KeeperRevoked(address keeper, bytes signature);

    event KeeperGranted(address keeper, address grantee);

    event KeeperDenied(address keeper, address grantee);

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

    function keeperCallAuthorized(
        address keeper,
        address caller
    ) external view returns (bool) {
        return true;
    }

    function getKeeperNotarized(
        address keeper,
        address notary
    ) public view returns (bool) {
        return keeperNotarizations[keeper][notary];
    }
}
