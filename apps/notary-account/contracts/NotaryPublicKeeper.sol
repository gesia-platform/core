// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicDAO.sol";
import "./NotaryKeeper.sol";

contract NotaryPublicKeeper is NotaryPublicDAO {
    mapping(address keeper => mapping(address notary => bool notarized))
        public keeperNotarizations;

    event KeeperNotarized(address keeper, bytes signature);

    event KeeperCanceled(address keeper, bytes signature);

    event KeeperGranted(address keeper, address grantee);

    event KeeperDenied(address keeper, address grantee);

    function notarizeKeeper(
        address keeper,
        bytes memory signature
    ) external onlyMember {
        require(
            NotaryKeeper(keeper).getGranter() == address(this),
            "keeper granter must be notary public."
        );

        address notary = getNotary(abi.encodePacked(keeper), signature);

        require(notary == msg.sender, "notraize must it self.");

        keeperNotarizations[keeper][notary] = true;

        emit KeeperNotarized(keeper, signature);
    }

    function cancelKeeper(address keeper, bytes memory signature) external {
        require(
            keeperNotarizations[keeper][msg.sender] == true,
            "sender has never notarized."
        );

        address notary = getNotary(abi.encodePacked(keeper), signature);

        require(notary == msg.sender, "notraize must it self.");

        delete keeperNotarizations[keeper][notary];

        emit KeeperCanceled(keeper, signature);
    }

    function grantKeeper(address keeper, address grantee) external onlyMember {
        require(
            keeperNotarizations[keeper][msg.sender] == true,
            "sender has never notarized."
        );

        NotaryKeeper notaryKeeper = NotaryKeeper(keeper);

        notaryKeeper.grant(grantee);

        emit KeeperGranted(address(notaryKeeper), grantee);
    }

    function denyKeeper(address keeper, address grantee) external onlyMember {
        require(
            keeperNotarizations[keeper][msg.sender] == true,
            "sender has never notarized."
        );

        NotaryKeeper notaryKeeper = NotaryKeeper(keeper);

        notaryKeeper.deny(grantee);

        emit KeeperDenied(address(notaryKeeper), grantee);
    }

    function getKeeperNotarized(
        address keeper,
        address notary
    ) public view returns (bool) {
        return keeperNotarizations[keeper][notary];
    }
}
