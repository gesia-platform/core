// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./AppPermissionBase.sol";
import "./AppPermissionNetworkAccess.sol";

contract AppPermission is
    Ownable,
    AppPermissionBase,
    AppPermissionNetworkAccess
{
    constructor(AppStore _appStore) {
        appStore = _appStore;
    }
}
