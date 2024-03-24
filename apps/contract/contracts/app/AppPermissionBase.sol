// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./AppStore.sol";

abstract contract AppPermissionBase {
    AppStore public appStore;
}
