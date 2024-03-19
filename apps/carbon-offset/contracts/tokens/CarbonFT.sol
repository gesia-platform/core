// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "../price/IPrice.sol";
import "./IMintOperator.sol";
import "../operator/IOperator.sol";
import "../fee/IFeeManager.sol";
import "../../../aaWallet/interfaces/ISignAuthorizer.sol";

contract Voucher1155Nft is
    ERC1155,
    IMintOperator,
    ReentrancyGuard,
    IPrice,
    Ownable
{
    using Strings for string;
    using SafeMath for uint256;
    using Counters for Counters.Counter;
    mapping(uint256 => uint256) public tokenSupply;
    mapping(bytes32 => bool) private transactionHashes;
    mapping(uint256 => mapping(address => uint256)) internal balances;
    mapping(uint256 => address) public creators;

    mapping(uint256 => string) tokenURIs;
    string public name;
    string public symbol;

    address public feeManager;

    uint256 private carbonPrice;

    mapping(uint256 => uint256) private carbonMapPrice;

    uint256 public minUSDTPrice = 1000000; // 1 USDT
    // token id counter incrementing by 1
    Counters.Counter private _tokenIdTracker;

    constructor(
        string memory _name,
        string memory _symbol,
        address _feeManager
    ) ERC1155("") {
        name = _name;
        symbol = _symbol;
        feeManager = _feeManager;
        // to start index from 1
        _tokenIdTracker.increment();
    }

    // for chainlink
    function mintByOperator(
        address _receiver,
        uint256 _amount,
        uint256 _tokenId,
        string memory _metadata
    ) external override onlyOwner {
        tokenURIs[_tokenId] = _metadata;
        tokenSupply[_tokenId] = tokenSupply[_tokenId].add(_amount);
        _mint(_receiver, _tokenId, _amount, "");
    }

    function mintBySignature(
        address _from,
        address _to,
        uint256 _amount,
        uint256 _tokenId,
        uint256 _nonce,
        string memory _metadata,
        bytes memory signature,
        uint256 _carbonPrice
    ) external onlyOwner {
        require(_carbonPrice >= minUSDTPrice, "price must be higher than min");
        require(
            _carbonPrice > carbonMapPrice[_tokenId],
            "price must be higher than old"
        );
        bytes32 hashMessage = keccak256(
            abi.encodePacked(_to, _tokenId, _amount, _nonce, address(this))
        );
        bytes32 hash = keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", hashMessage)
        );
        address signer = recoverSigner(hash, signature);
        require(
            ISignAuthorizer(_from).isAuthorizedSigner(signer),
            "Signature does not match the sender"
        );
        require(
            !transactionHashes[hashMessage],
            "Transaction already processed"
        );
        transactionHashes[hashMessage] = true;

        // company fee amount
        uint256 calculatedAmount = IFeeManager(feeManager).feeAmount(_amount);
        uint256 remainAmount = _amount.sub(calculatedAmount);

        tokenURIs[_tokenId] = _metadata;
        tokenSupply[_tokenId] = tokenSupply[_tokenId].add(_amount);
        _mint(_to, _tokenId, remainAmount, "");
        _mint(
            IFeeManager(feeManager).feeAddress(),
            _tokenId,
            calculatedAmount,
            ""
        );
        carbonMapPrice[_tokenId] = _carbonPrice;
    }

    function transferBySignature(
        address from,
        address to,
        uint256 tokenId,
        uint256 amount,
        uint256 nonce,
        bytes memory signature
    ) external nonReentrant onlyOwner {
        bytes32 hashMessage = keccak256(
            abi.encodePacked(from, to, tokenId, amount, nonce, address(this))
        );
        bytes32 hash = keccak256(
            abi.encodePacked("\x19Ethereum Signed Message:\n32", hashMessage)
        );
        address signer = recoverSigner(hash, signature);
        require(
            ISignAuthorizer(from).isAuthorizedSigner(signer),
            "Signature does not match the sender"
        );
        require(
            !transactionHashes[hashMessage],
            "Transaction already processed"
        );
        transactionHashes[hashMessage] = true;

        // company fee amount
        uint256 calculatedAmount = IFeeManager(feeManager).feeAmount(amount);
        uint256 remainAmount = amount.sub(calculatedAmount);

        _safeTransferFrom(from, to, tokenId, remainAmount, "");
        _safeTransferFrom(
            from,
            IFeeManager(feeManager).feeAddress(),
            tokenId,
            calculatedAmount,
            ""
        );
    }

    function uri(uint256 _id) public view override returns (string memory) {
        require(_exists(_id), "MultiCollection#uri: NONEXISTENT_TOKEN");
        return tokenURIs[_id];
    }

    function _exists(uint256 _id) internal view returns (bool) {
        return creators[_id] != address(0);
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

    function getCarbonPrice(
        uint256 tokenId
    ) external view override returns (uint256) {
        return carbonMapPrice[tokenId];
    }
}
