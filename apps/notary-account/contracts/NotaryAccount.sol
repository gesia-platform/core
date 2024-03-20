// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

contract NotaryAccount {
    string public domain; // notarized domain
    address public owner; // account owner

    address public immutable notaryPublic;

    mapping(address => string) public notaryURLs;

    modifier onlyNotaryPublic() {
        require(
            notaryPublic == msg.sender,
            "can only be called by notary public."
        );
        _;
    }

    constructor(address _notaryPublic, string memory _domain, address _owner) {
        notaryPublic = _notaryPublic;
        domain = _domain;
        owner = _owner;
    }

    function notarize(
        address notary,
        string memory url
    ) external onlyNotaryPublic {
        notaryURLs[notary] = url;
    }

    function revoke(address notary) external onlyNotaryPublic {
        delete notaryURLs[notary];
    }

    function setOwner(address _owner) external onlyNotaryPublic {
        owner = _owner;
    }

    function getNotaryURL(address notary) public view returns (string memory) {
        return notaryURLs[notary];
    }

    function getDomain() public view returns (string memory) {
        return domain;
    }

    function getOwner() public view returns (address) {
        return owner;
    }

    function getNotaryPublic() public view returns (address) {
        return notaryPublic;
    }
}
