// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./Carbon1155Nft.sol";

contract Carbon1155NftFactory {

    address public operatorManager;
    address public feeManager;
    address public voucherAddress;
    mapping(bytes32 => bool) private transactionHashes;

    event CreateCollection(address from, uint256 tokenId, string name, string symbol, address collection);

    constructor(address _operatorManager, address _feeManager, address _voucherAddress){
        operatorManager = _operatorManager;
        voucherAddress = _voucherAddress;
        feeManager = _feeManager;
    }

    function createCarbon1155NFT(uint256 tokenId, string memory name, string memory symbol) external {
        Carbon1155Nft collection = new Carbon1155Nft(name, symbol, operatorManager, feeManager, voucherAddress, tokenId);
        emit CreateCollection(msg.sender, tokenId, name, symbol, address(collection));
    }

    function recoverSigner(
        bytes32 _ethSignedMessageHash,
        bytes memory _signature
    ) internal pure returns (address) {
        (bytes32 r, bytes32 s, uint8 v) = splitSignature(_signature);

        return ecrecover(_ethSignedMessageHash, v, r, s);
    }

    function splitSignature(
        bytes memory sig
    ) internal pure returns (bytes32 r, bytes32 s, uint8 v) {
        require(sig.length == 65, "invalid signature length");
        assembly {
            r := mload(add(sig, 32))
            s := mload(add(sig, 64))
            v := byte(0, mload(add(sig, 96)))
        }
    }

}
