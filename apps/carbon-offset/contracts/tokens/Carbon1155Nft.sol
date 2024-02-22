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
import "../operator/IOperator.sol";
import "../fee/IFeeManager.sol";
import "../price/IPrice.sol";
import "../../../aaWallet/interfaces/ISignAuthorizer.sol";

contract Carbon1155Nft is ERC1155, ReentrancyGuard, IPrice {
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

    mapping(uint256 => uint256) private _tokenWeight;
    address public voucherAddress;
    uint256 public voucherTokenId;
    address public operatorManager;
    address public feeManager;
    uint256 public feeAmount = 100; // 1%
    uint256 public minUSDTPrice = 1000000; // 1 USDT

    // token id counter incrementing by 1
    Counters.Counter private _tokenIdTracker;
    uint256 private carbonPrice;
    mapping(uint256 => uint256) private carbonMapPrice;

    constructor(
        string memory _name,
        string memory _symbol,
        address _operatorManager,
        address _feeManager,
        address _voucherAddress,
        uint256 _tokenId
    ) ERC1155("") {
        name = _name;
        symbol = _symbol;
        voucherAddress = _voucherAddress;
        voucherTokenId = _tokenId;
        operatorManager = _operatorManager;
        feeManager = _feeManager;
        // to start index from 1
        _tokenIdTracker.increment();
    }

    modifier operatorsOnly() {
        require(IOperator(operatorManager).isOperator(msg.sender), "#operatorsOnly:");
        _;
    }

    function mint(
        address _from,
        uint256 _carbonAmount,
        uint256 _nonce,
        string memory _metadata,
        bytes memory _signature,
        uint256 _carbonPrice
    ) external operatorsOnly {
        require(_carbonAmount > 0, "must be higher than zero");
        require(_carbonPrice >= minUSDTPrice, "price must be higher than min");
        require(_carbonPrice > carbonMapPrice[_tokenIdTracker.current()], "price must be higher than old");

        bytes32 hashMessage = keccak256(abi.encodePacked(_from, _carbonAmount, _nonce, address(this)));
        bytes32 hash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hashMessage));
        address signer = recoverSigner(hash, _signature);
        require(ISignAuthorizer(_from).isAuthorizedSigner(signer), "Signature does not match the sender");
        require(!transactionHashes[hashMessage], "Transaction already processed");
        transactionHashes[hashMessage] = true;

        require(IERC1155(voucherAddress).balanceOf(_from, voucherTokenId) >= _carbonAmount, "lack of carbon balance");
        // company fee amount
        uint256 calculatedAmount = IFeeManager(feeManager).feeAmount(_carbonAmount);
        uint256 remainAmount = _carbonAmount.sub(calculatedAmount);
        // transfer fee to company owner
        IERC1155(voucherAddress).safeTransferFrom(_from, IFeeManager(feeManager).feeAddress(), voucherTokenId, calculatedAmount, "");

        creators[_tokenIdTracker.current()] = _from;
        tokenURIs[_tokenIdTracker.current()] = _metadata;
        tokenSupply[_tokenIdTracker.current()] = remainAmount;
        _mint(_from, _tokenIdTracker.current(), remainAmount, "");
        _tokenWeight[_tokenIdTracker.current()] = remainAmount;
        carbonMapPrice[_tokenIdTracker.current()] = _carbonPrice;
        _tokenIdTracker.increment();
    }

    function transferWithSignature(
        address from,
        address to,
        uint256 tokenId,
        uint256 amount,
        uint256 nonce,
        bytes memory signature
    ) external nonReentrant operatorsOnly {
        require(msg.sender == to, "Only receiver call");
        bytes32 hashMessage = keccak256(abi.encodePacked(from, to, tokenId, amount, nonce, address(this)));
        bytes32 hash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hashMessage));
        address signer = recoverSigner(hash, signature);
        require(ISignAuthorizer(from).isAuthorizedSigner(signer), "Signature does not match the sender");
        require(!transactionHashes[hashMessage], "Transaction already processed");
        transactionHashes[hashMessage] = true;

        require(IERC1155(voucherAddress).balanceOf(from, voucherTokenId) >= amount, "lack of carbon balance");
        // company fee amount
        uint256 calculatedAmount = IFeeManager(feeManager).feeAmount(amount);
        uint256 remainAmount = amount.sub(calculatedAmount);
        // transfer fee to company owner
        IERC1155(voucherAddress).safeTransferFrom(from, IFeeManager(feeManager).feeAddress(), voucherTokenId, calculatedAmount, "");
        _safeTransferFrom(from, to, tokenId, remainAmount, "");
    }

    function transferWithGift(
        address from,
        address to,
        uint256 tokenId,
        uint256 amount,
        uint256 nonce,
        bytes memory signature
    ) external nonReentrant {
        require(msg.sender == to, "Only receiver call");
        bytes32 hashMessage = keccak256(abi.encodePacked(from, to, tokenId, amount, nonce, address(this)));
        bytes32 hash = keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hashMessage));
        address signer = recoverSigner(hash, signature);
        require(ISignAuthorizer(from).isAuthorizedSigner(signer), "Signature does not match the sender");
        require(!transactionHashes[hashMessage], "Transaction already processed");
        transactionHashes[hashMessage] = true;
        _safeTransferFrom(from, to, tokenId, amount, "");
    }

    function uri(
        uint256 _id
    ) override public view returns (string memory) {
        require(_exists(_id), "MultiCollection#uri: NONEXISTENT_TOKEN");
        return tokenURIs[_id];
    }

    function _exists(
        uint256 _id
    ) internal view returns (bool) {
        return creators[_id] != address(0);
    }


    function tokenWeight(uint256 tokenId) public view returns (uint256) {
        require(_exists(tokenId), "ERC721Metadata: URI query for nonexistent token");
        return _tokenWeight[tokenId];
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

    function getCarbonPrice(uint256 tokenId) external view override returns (uint256) {
        return carbonMapPrice[tokenId];
    }

}
