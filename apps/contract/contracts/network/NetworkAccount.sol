// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";

contract NetworkAccount is Ownable {
    string public name;
    string public gatewayURL;

    constructor(string memory _name, string memory _gatewayURL) {
        name = _name;
        gatewayURL = _gatewayURL;
    }

    function setGatewayURL(string memory _gatewayURL) public onlyOwner {
        gatewayURL = _gatewayURL;
    }

    function getName() public view returns (string memory) {
        return name;
    }

    function getGatewayURL() public view returns (string memory) {
        return gatewayURL;
    }
}
