// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicAccount.sol";
import "./NotaryPublicKeeper.sol";
import "./NotaryPublicModule.sol";

contract NotaryPublic is
    NotaryPublicAccount,
    NotaryPublicModule,
    NotaryPublicKeeper
{
    function moduleCallAuthorized(
        address module,
        address sender
    ) public view returns (bool) {
        return true;
    }
}
