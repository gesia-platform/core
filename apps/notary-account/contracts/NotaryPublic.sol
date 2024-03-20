// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicAccount.sol";
import "./NotaryPublicKeeper.sol";

contract NotaryPublic is NotaryPublicAccount, NotaryPublicKeeper {}
