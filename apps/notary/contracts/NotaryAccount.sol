// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

contract NotaryAccount {
    string public domain;
    uint256 public applicationID;

    constructor(string memory _domain, uint256 _applicationID) {
        domain = _domain;
        applicationID = _applicationID;
    }

    function getDomain() public view returns (string memory) {
        return domain;
    }

    function getApplicationID() public view returns (uint256) {
        return applicationID;
    }
}
