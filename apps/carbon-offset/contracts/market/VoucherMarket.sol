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
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@chainlink/contracts/src/v0.8/ConfirmedOwner.sol";

contract VoucherMarket is ERC1155Holder{
    using Counters for Counters.Counter;
    using SafeMath for uint256;
    using SafeERC20 for ERC20;

    struct MarketItem {
        address voucherContract; // voucher contract address
        uint256 tokenId; // voucher tokenId
        uint256 amount; // voucher amount
        uint256 price; // price (USDT (decimals 6))
        address seller; // seller wallet address
    }

    mapping(uint256 => MarketItem) public _marketItemMap;
    mapping(address => bool) voucherContractMap;
    Counters.Counter private _marketItemIds;
    address public usdtContractAddress;
    address public operatorManager;
    address public whitelistManager;
    address public feeManager;
    address public minterContract; // minter contract
    uint256 public minVoucherAmount = 1; // 1 Token

    event VerificationVoucherContract(
        address indexed voucherContract,
        bool isVerified
    );

    event MinPriceChange(
        uint256 price,
        bool isBnb
    );

    event TokenPlaced(
        address indexed voucherContract,
        uint256 indexed marketId,
        uint256 tokenId,
        uint256 amount,
        address seller,
        uint256 price);

    event TokenUnPlaced(
        address indexed voucherContract,
        uint256 indexed marketId,
        uint256 tokenId,
        uint256 deductedAmount,
        uint256 remainAmount,
        address seller,
        uint256 price);

    event TokenSold(
        address indexed voucherContract,
        uint256 indexed marketId,
        uint256 tokenId,
        uint256 amount,
        address buyer,
        address seller,
        uint256 price,
        uint256 totalPrice,
        uint256 feeAmount, // amount to fee address
        uint256 remainAmount);

    constructor(
        address _usdtContractAddress,
        address _operatorManager,
        address _minterContract,
        address _whitelistManager,
        address _feeManager){
        usdtContractAddress = _usdtContractAddress;
        operatorManager = _operatorManager;
        minterContract = _minterContract;
        whitelistManager = _whitelistManager;
        feeManager = _feeManager;
    }

    receive() external payable {
    }

    modifier operatorsOnly() {
        require(IOperator(operatorManager).isOperator(msg.sender), "#operatorsOnly:");
        _;
    }

    function verifyVoucherContract(address _voucherContract) external operatorsOnly {
        voucherContractMap[_voucherContract] = true;
        emit VerificationVoucherContract(_voucherContract, true);
    }

    function unVerifyVoucherContract(address _voucherContract) external operatorsOnly {
        voucherContractMap[_voucherContract] = false;
        emit VerificationVoucherContract(_voucherContract, false);
    }

    // place voucher token to sell
    // _amount : amount of voucher token
    // _tokenId : tokenId of voucher token
    // _voucherContract : contract address of voucher token
    // _perTokenPrice : 1 vocher of token if isBnB decimal 18 if not Bnb decimal in 6 (for usdt)

    //  if (allowance) {
    //   place()
    //  else {
    //   approve();
    //   place
    //  }
    function place(uint256 _tokenId, uint256 _amount, address _voucherContract, uint256 _perTokenPrice) external {
        // check voucher contract verification
        require(voucherContractMap[_voucherContract], "Not Valid Voucher Contract");
        // check for minimum amount
        require(_amount >= minVoucherAmount, "Minimum purchase amount error");
        // check for whitelist
        require(IWhitelist(whitelistManager).isWhitelist(_voucherContract,_tokenId, msg.sender), "not in whitelist");
        // check min _perTokenPrice
        require(_perTokenPrice >= IPrice(_voucherContract).getCarbonPrice(_tokenId), "minimum usdt price issue");
        // increment marketId
        _marketItemIds.increment();
        uint256 marketId = _marketItemIds.current();
        // create market map for each marketId
        _marketItemMap[marketId] = MarketItem(
            _voucherContract,
            _tokenId,
            _amount,
            _perTokenPrice,
            address(msg.sender)
        );
        // voucher token : transfer from seller wallet to contract
        IERC1155(_voucherContract).safeTransferFrom(msg.sender, address(this), _tokenId, _amount, "");
        // trigger event
        emit TokenPlaced(_voucherContract, marketId, _tokenId, _amount, address(msg.sender), _perTokenPrice);
    }

    // unPlance : cancel voucher token amount
    // _marketId : marketId which is generated from place function
    // _amount : voucher amount to deduct (decimals 18)
    function unPlace(uint256 _marketId, uint256 _amount) external {
        // check amount must be higher than 0
        require(_amount > 0, "Must be higher than zero");
        MarketItem storage marketItem = _marketItemMap[_marketId];
        // check owner
        require(marketItem.seller == address(msg.sender) || IOperator(operatorManager).isOperator(address(msg.sender)), "Not ownerOf or Operators");
        // check market amount to deduct
        require(marketItem.amount >= _amount, "Not Enough amount");
        // deduct amount from market amount
        marketItem.amount = marketItem.amount.sub(_amount);
        // send deducted amount to seller address
        IERC1155(marketItem.voucherContract).safeTransferFrom(address(this), msg.sender, marketItem.tokenId, _amount, "");
        // trigger event
        emit TokenUnPlaced(marketItem.voucherContract, _marketId, marketItem.tokenId, _amount, marketItem.amount, marketItem.seller, marketItem.price);
    }

    // purchaseInUSDT : voucher token purchasable by USDT
    // _marketId : marketId which is generated from place function
    // _amount : voucher amount to buy

    //  if (allowance) {
    //   purchaseInUSDT()
    //  else {
    //   approve();
    //   purchaseInUSDT
    //  }
    //
    function purchaseInUSDT(uint256 _marketId, uint256 _amount) external {
        // check minimumAmount
        require(_amount >= minVoucherAmount, "Minimum purchase amount error");
        // get marketItem
        MarketItem storage marketItem = _marketItemMap[_marketId];
        // check for whitelist
        require(IWhitelist(whitelistManager).isWhitelist(marketItem.voucherContract, marketItem.tokenId , msg.sender), "not in whitelist");
        // check amount from market amount
        require(marketItem.amount >= _amount, "Not Enough amount");
        // divide totalPrice to decimal
        uint256 totalPrice = marketItem.price.mul(_amount);
        // fee amount
        uint256 feeAmount = IFeeManager(feeManager).feeAmount(totalPrice);
        // remain amount
        uint256 remainAmount = totalPrice.sub(feeAmount);
        // check contract balance
        require(IERC1155(marketItem.voucherContract).balanceOf(address(this), marketItem.tokenId) >= _amount, "Insufficient tokens in contract");
        // check user's wallet balance
        require(IERC20(usdtContractAddress).balanceOf(msg.sender) >= totalPrice, "Insufficient USDT balance");
        // deduct amount
        marketItem.amount = marketItem.amount.sub(_amount);
        // transfer remainAmount to seller
        ERC20(usdtContractAddress).safeTransferFrom(msg.sender, marketItem.seller, remainAmount);
        // transfer feeAmount to feeAddress
        ERC20(usdtContractAddress).safeTransferFrom(msg.sender, IFeeManager(feeManager).feeAddress(), feeAmount);
        // transfer voucher
        IERC1155(marketItem.voucherContract).safeTransferFrom(address(this), msg.sender, marketItem.tokenId, _amount, "");
        // trigger event
        emit TokenSold(marketItem.voucherContract, _marketId, marketItem.tokenId, _amount, address(msg.sender), marketItem.seller, marketItem.price, totalPrice, feeAmount, remainAmount);
    }
}
