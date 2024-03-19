// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract NotaryAccount {
    using ECDSA for bytes32;

    string public chain;
    string public domain;

    uint public verificationBlockHeight;

    address public immutable notaryPublic;

    mapping(address => string) public notarizedUrls;

    event RPCNotarized(address notary, string url);

    event RPCCanceled(address notary);

    constructor(
        address _notaryPublic,
        uint _verificationBlockHeight,
        string memory _chain,
        string memory _domain
    ) {
        notaryPublic = _notaryPublic;
        verificationBlockHeight = _verificationBlockHeight;
        chain = _chain;
        domain = _domain;
    }

    function cancelRPC(bytes memory signature) external {
        address notary = keccak256(abi.encodePacked(domain))
            .toEthSignedMessageHash()
            .recover(signature);

        delete notarizedUrls[notary];

        emit RPCCanceled(notary);
    }

    function notarizeRPC(bytes memory signature, string memory url) external {
        require(
            notaryPublic == msg.sender,
            "can only be called by notary public."
        );

        address notary = keccak256(abi.encodePacked(domain))
            .toEthSignedMessageHash()
            .recover(signature);

        notarizedUrls[notary] = url;

        emit RPCNotarized(notary, url);
    }

    function getNotarizedUrl(
        address notary
    ) public view returns (string memory) {
        return notarizedUrls[notary];
    }

    function getDomain() public view returns (string memory) {
        return domain;
    }

    function getChain() public view returns (string memory) {
        return chain;
    }

    function getVerificationBlockHeight() public view returns (uint) {
        return verificationBlockHeight;
    }
}
