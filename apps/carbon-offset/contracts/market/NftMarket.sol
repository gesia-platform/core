// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../operator/IOperator.sol";
import "../whitelist/IWhitelist.sol";
import "../fee/IFeeManager.sol";
import "../price/IPrice.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";


contract NftMarket is ERC1155Holder {
    using Counters for Counters.Counter;
    using SafeMath for uint256;
    using SafeERC20 for ERC20;

    struct NFTMarketItem {
        address contractAddress; // nft contract address
        uint256 tokenId; // nft tokenId
        uint256 amount; // nft amount
        uint256 price;  // price (USDT (decimals 6) or BNB (decimals 18))
        address seller; // seller wallet address
    }

    mapping(uint256 => NFTMarketItem) public _marketItemMap;
    mapping(address => bool) nftContractMap;
    Counters.Counter private _marketItemIds;
    address public usdtContractAddress;
    address public operatorManager;
    address public whitelistManager;
    address public feeManager;
    bool public isWhitelistEnabled;


    event MinPriceChange(
        uint256 price,
        bool isBnb
    );

    event VerificationNftContract(
        address indexed nftContract,
        bool isVerified
    );

    event TokenPlaced(
        address indexed nftContract,
        uint256 indexed tokenId,
        uint256 indexed marketId,
        uint256 amount,
        address seller,
        uint256 price);

    event TokenUnPlaced(
        address indexed nftContract,
        uint256 indexed tokenId,
        uint256 indexed marketId,
        uint256 deductedAmount,
        uint256 remainAmount,
        address seller,
        uint256 price);

    event TokenSold(
        address indexed nftContract,
        uint256 indexed tokenId,
        uint256 indexed marketId,
        uint256 amount,
        address buyer,
        address seller,
        uint256 price,
        uint256 totalPrice,
        uint256 feeAmount, // fee amount to fee address
        uint256 remainAmount);

    constructor(
        address _usdtContractAddress,
        address _whitelistManager,
        address _operatorManager,
        address _feeManager){
        usdtContractAddress = _usdtContractAddress;
        whitelistManager = _whitelistManager;
        operatorManager = _operatorManager;
        feeManager = _feeManager;
    }

    receive() external payable {
    }

    modifier operatorsOnly() {
        require(IOperator(operatorManager).isOperator(msg.sender), "#operatorsOnly:");
        _;
    }

    function changeWhitelistStatus(bool _status) external operatorsOnly {
        isWhitelistEnabled = _status;
    }

    function verifyNftContract(address _nftContract) external operatorsOnly {
        nftContractMap[_nftContract] = true;
        emit VerificationNftContract(_nftContract, true);
    }


    function unVerifyNftContract(address _nftContract) external operatorsOnly {
        nftContractMap[_nftContract] = false;
        emit VerificationNftContract(_nftContract, false);
    }

    // place voucher token to sell
    // _amount : amount
    // _nftContract : contract address of nft
    // _tokenId : nft tokenId
    // _perNftPrice : 1 nft if isBnB decimal 18 if not Bnb decimal in 6 (for usdt)
    // _isBnb : if true -> BNB, if not true -> USDT

    // if(isApproval()) {
    //  place();
    // } else {
    //  setApprovalForAll();
    //  place();
    // }
    // setApprovalForAll (_nftMarketContract, true);
    function place(uint256 _amount, address _nftContract, uint256 _tokenId, uint256 _perNftPrice) external {
        // check nft contract verification
        require(nftContractMap[_nftContract], "Not Valid Nft Contract");
        // check amount must be higher than 0
        require(_amount > 0, "Must be higher than zero");
        require(_perNftPrice >= IPrice(_nftContract).getCarbonPrice(_tokenId), "min carbon price issue");
        // check for whitelist
        if (isWhitelistEnabled) {
            require(IWhitelist(whitelistManager).isWhitelist(_nftContract, _tokenId, msg.sender), "not in whitelist");
        }
        // increment marketId
        _marketItemIds.increment();
        uint256 marketId = _marketItemIds.current();
        // create market map for each marketId
        _marketItemMap[marketId] = NFTMarketItem(
            _nftContract,
            _tokenId,
            _amount,
            _perNftPrice,
            address(msg.sender)
        );
        // nft : transfer from seller wallet to contract
        IERC1155(_nftContract).safeTransferFrom(msg.sender, address(this), _tokenId, _amount, "");
        // trigger event
        emit TokenPlaced(_nftContract, _tokenId, marketId, _amount, address(msg.sender), _perNftPrice);
    }

    // unPlace : cancel nft token amount
    // _marketId : marketId which is generated from place function
    // _amount : nft amount to deduct
    function unPlace(uint256 _marketId, uint256 _amount) external {
        // check amount must be higher than 0
        require(_amount > 0, "Must be higher than zero");
        // get marketItem
        NFTMarketItem storage marketItem = _marketItemMap[_marketId];
        // check owner
        require(marketItem.seller == address(msg.sender) || IOperator(operatorManager).isOperator(address(msg.sender)), "Not ownerOf or Operators");
        // check market amount to deduct
        require(marketItem.amount >= _amount, "Not Enough amount");
        // deduct amount from market amount
        marketItem.amount = marketItem.amount.sub(_amount);
        // send deducted amount to seller address
        IERC1155(marketItem.contractAddress).safeTransferFrom(address(this), msg.sender, marketItem.tokenId, _amount, "");
        // trigger event
        emit TokenUnPlaced(marketItem.contractAddress, marketItem.tokenId, _marketId, _amount, marketItem.amount, marketItem.seller, marketItem.price);
    }

    // purchaseInUSDT : voucher token purchasable by USDT
    // _marketId : marketId which is generated from place function
    // _amount : voucher amount to buy (decimals 18)

    //  if (allowance) {
    //   purchaseInUSDT()
    //  else {
    //   approve();
    //   purchaseInUSDT
    //  }
    function purchaseInUSDT(uint256 _marketId, uint256 _amount) external {
        // check amount must be higher than 0
        require(_amount > 0, "Must be higher than zero");
        // get marketItem
        NFTMarketItem storage marketItem = _marketItemMap[_marketId];
        // check amount from market amount
        require(marketItem.amount >= _amount, "Not Enough amount");
        // check for whitelist
        if (isWhitelistEnabled) {
            require(IWhitelist(whitelistManager).isWhitelist(marketItem.contractAddress, marketItem.tokenId, msg.sender), "not in whitelist");
        }
        // divide totalPrice to decimal
        uint256 totalPrice = marketItem.price.mul(_amount);
        // fee amount
        uint256 feeAmount = IFeeManager(feeManager).feeAmount(totalPrice);
        // remain amount
        uint256 remainAmount = totalPrice.sub(feeAmount);
        // check usdt amount in balance
        require(ERC20(usdtContractAddress).balanceOf(msg.sender) >= totalPrice, "Lack Of USDT");
        marketItem.amount = marketItem.amount.sub(_amount);
        // transfer remainAmount to seller
        ERC20(usdtContractAddress).safeTransferFrom(msg.sender, marketItem.seller, remainAmount);
        // transfer feeAmount to feeAddress
        ERC20(usdtContractAddress).safeTransferFrom(msg.sender, IFeeManager(feeManager).feeAddress(), feeAmount);
        // transfer voucher token to buyer
        IERC1155(marketItem.contractAddress).safeTransferFrom(address(this), msg.sender, marketItem.tokenId, _amount, "");
        // trigger event
        emit TokenSold(marketItem.contractAddress, marketItem.tokenId, _marketId, _amount, address(msg.sender), marketItem.seller, marketItem.price, totalPrice, feeAmount, remainAmount);
    }
}
