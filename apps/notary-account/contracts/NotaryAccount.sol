// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract NotaryAccount {
    using ECDSA for bytes32;

    string public domain; // notarized domain
    address public caller; // notarized caller

    address public immutable notaryPublic;

    mapping(address => string) public notaryURLs;

    event Notarized(address notary);

    event Canceled(address notary);

    modifier onlyNotaryPublic() {
        require(
            notaryPublic == msg.sender,
            "can only be called by notary public."
        );
        _;
    }

    constructor(address _notaryPublic, string memory _domain, address _caller) {
        notaryPublic = _notaryPublic;
        domain = _domain;
        caller = _caller;
    }

    function notarize(
        bytes memory signature,
        string memory url
    ) external onlyNotaryPublic {
        address notary = keccak256(abi.encodePacked(domain, caller))
            .toEthSignedMessageHash()
            .recover(signature);

        notaryURLs[notary] = url;

        emit Notarized(notary);
    }

    function setCaller(address _caller) external onlyNotaryPublic {
        caller = _caller;
    }

    function cancel(bytes memory signature) external {
        address notary = keccak256(abi.encodePacked(domain, caller))
            .toEthSignedMessageHash()
            .recover(signature);

        delete notaryURLs[notary];

        emit Canceled(notary);
    }

    function getNotaryURL(address notary) public view returns (string memory) {
        return notaryURLs[notary];
    }

    function getDomain() public view returns (string memory) {
        return domain;
    }

    function getCaller() public view returns (address) {
        return caller;
    }
}
