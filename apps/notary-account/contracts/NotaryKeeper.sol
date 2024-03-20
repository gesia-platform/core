// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./NotaryAccount.sol";

abstract contract NotaryKeeper {
    mapping(address => bool) public grantees;

    address public immutable granter;

    modifier onlyGranter() {
        require(msg.sender == granter, "can only be called by granter.");
        _;
    }

    modifier onlyGrantee() {
        require(grantees[msg.sender] == true, "can only be called by grantee.");
        _;
    }

    constructor(address _granter) {
        granter = _granter;
    }

    function grant(address grantee) external onlyGranter {
        grantees[grantee] = true;
    }

    function deny(address grantee) external onlyGranter {
        delete grantees[grantee];
    }

    function getGranter() external view returns (address) {
        return granter;
    }
}
