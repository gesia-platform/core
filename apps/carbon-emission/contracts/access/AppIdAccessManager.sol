// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";

contract AppIdAccessManager is Ownable {
    mapping(uint256 => mapping(address => bool)) private _authorizedAddresses;

    event AuthorizationUpdated(uint256 indexed appId, address indexed _address, bool isAuthorized);

    function isAuthorized(uint256 appId, address _address) public view returns (bool) {
        return _authorizedAddresses[appId][_address];
    }

    function authorizeAddress(uint256 appId, address _address) public onlyOwner {
        _authorizedAddresses[appId][_address] = true;
        emit AuthorizationUpdated(appId, _address, true);
    }

    function removeAuthorization(uint256 appId, address _address) public onlyOwner {
        _authorizedAddresses[appId][_address] = false;
        emit AuthorizationUpdated(appId, _address, false);
    }
}
