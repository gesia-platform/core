// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryAccount.sol";
import "./NotaryPublicDAO.sol";

contract NotaryPublicModule is NotaryPublicDAO {
    function moduleCallAuthorized(
        address module,
        address sender
    ) public view returns (bool) {
        return true;
    }
}
