// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

contract NotaryAccount {
    string public domain;

    address public owner;

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

    function setNotaryURL(
        address notary,
        string memory url
    ) external onlyNotaryPublic {
        if (bytes(url).length == 0) {
            delete notaryURLs[notary];
        } else {
            notaryURLs[notary] = url;
        }
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
