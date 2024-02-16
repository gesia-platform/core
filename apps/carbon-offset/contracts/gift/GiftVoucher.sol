// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155Receiver.sol";
import "@openzeppelin/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/security/ReentrancyGuard.sol";

contract CarbonNFT is ERC1155,  ReentrancyGuard, Ownable {
    using Strings for string;
    using SafeMath for uint256;
    using Counters for Counters.Counter;
    // token id counter incrementing by 1
    Counters.Counter private _tokenIdTracker;
    mapping(uint256 => uint256) public tokenSupply;
    mapping(bytes32 => bool) private transactionHashes;
    mapping (uint256 => mapping(address => uint256)) internal balances;

    mapping(uint256 => address) public creators;

    mapping(uint256 => string) tokenURIs;
    string public name;
    string public symbol;

    mapping(uint256 => uint256) private _tokenWeight;
    address public nvmToken;

    constructor(
        string memory _name,
        string memory _symbol,
        address _nvmToken
    ) ERC1155("") {
        name = _name;
        symbol = _symbol;
        nvmToken = _nvmToken;
        // to start index from 1
        _tokenIdTracker.increment();
    }

    function mint(
        address _to,
        uint256 _carbonAmount, // 100
        string memory _metadata // url
    ) external onlyOwner {
        require(_carbonAmount > 0, "must be higher than zero");
        require(IERC20(nvmToken).balanceOf(msg.sender) >= _carbonAmount, "lack of carbon balance");
        require(!_exists(_tokenIdTracker.current()), "token _id already exists");
        IERC20(nvmToken).transferFrom(msg.sender, address(this), _carbonAmount);
        creators[_tokenIdTracker.current()] = _to;
        tokenURIs[_tokenIdTracker.current()] = _metadata;
        tokenSupply[_tokenIdTracker.current()] = _carbonAmount.div(10 ** 18);
        _mint(_to, _tokenIdTracker.current(), _carbonAmount.div(10 ** 18), "");
        _tokenWeight[_tokenIdTracker.current()] = _carbonAmount.div(10 ** 18);
        _tokenIdTracker.increment();
    }

    function transferWithSignature(
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
        require(signer == from, "Signature does not match the sender");
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
}
