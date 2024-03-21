// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicAccount.sol";
import "./NotaryPublicModule.sol";
import "./NotaryKeeper.sol";

contract NotaryPublic is NotaryPublicAccount, NotaryPublicModule {
    mapping(address keeper => bool authorized) public keeperAuthorizations;

    function authorizeKeeper(address keeper) external {
        require(
            keeperAuthorizations[keeper] == false,
            "keeper is already authorized."
        );

        // 1. 심사된 모듈이여야 함.
        require(
            moduleAuditResults[NotaryKeeper(keeper).notaryModuleAddress()] ==
                true,
            "module is not audited"
        );

        keeperAuthorizations[keeper] = true;
    }

    function getKeeperAuthorized(address keeper) public view returns (bool) {
        return keeperAuthorizations[keeper];
    }

    function getKeeperAuthorizedApplication(
        address keeper,
        uint256 applicationID
    ) public view returns (bool) {
        NotaryAccount account = NotaryAccount(
            NotaryKeeper(keeper).notaryAccountAddress()
        );

        return account.getApplicationID() == applicationID;
    }
}
