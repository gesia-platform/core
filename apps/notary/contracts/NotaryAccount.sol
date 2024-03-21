// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

contract NotaryAccount {
    string public domain;
    uint256 public applicationID;
    address public owner;

    address public immutable notaryPublic;

    modifier onlyNotaryPublic() {
        require(msg.sender == notaryPublic, "sender is not notary public.");
        _;
    }

    constructor(
        address _notaryPublic,
        string memory _domain,
        uint256 _applicationID,
        address _owner
    ) {
        notaryPublic = _notaryPublic;
        domain = _domain;
        applicationID = _applicationID;
        owner = _owner;
    }

    function setOwner(address _owner) external onlyNotaryPublic {
        owner = _owner;
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

    function getApplicationID() public view returns (uint256) {
        return applicationID;
    }
}
