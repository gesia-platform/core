// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../operator/IOperator.sol";
import "../fee/IFeeManager.sol";
import "../price/IPrice.sol";
import "../openzeppelin//utils/math/SafeMath.sol";
import "../openzeppelin//token/ERC20/utils/SafeERC20.sol";
import "../openzeppelin//token/ERC20/ERC20.sol";
import "../openzeppelin//utils/Counters.sol";
import "../openzeppelin//token/ERC1155/ERC1155.sol";
import "../openzeppelin//token/ERC1155/IERC1155.sol";
import "../openzeppelin//token/ERC1155/utils/ERC1155Holder.sol";

contract MMarket is ERC1155Holder {
    using Counters for Counters.Counter;
    using SafeMath for uint256;
    using SafeERC20 for ERC20;

    struct MarketItem {
        uint256 tokenId; // voucher contract address
        address voucherNftContract; // voucher nft contract address
        uint256 amount; // voucher amount in decimals 18
        uint256 price; // price USD
        address seller; // seller wallet address
    }

    mapping(uint256 => MarketItem) public _marketItemMap;
    mapping(address => bool) voucherNftContractMap;
    Counters.Counter private _marketItemIds;

    address public operatorManager;
    address public whitelistManager;
    address public feeManager;
    mapping(bytes32 => bool) private transactionHashes;

    receive() external payable {
    }

    modifier operatorsOnly() {
        require(IOperator(operatorManager).isOperator(msg.sender), "#operatorsOnly:");
        _;
    }

    event VerificationVoucherNftContract(
        address indexed voucherNftContract,
        bool isVerified
    );

    event TokenPlaced(
        address indexed voucherNftContract,
        uint256 indexed tokenId,
        uint256 indexed marketId,
        uint256 amount,
        address seller,
        uint256 price);

    event TokenUnPlaced(
        address indexed voucherNftContract,
        uint256 indexed tokenId,
        uint256 indexed marketId,
        uint256 deductedAmount,
        uint256 remainAmount,
        address seller,
        uint256 price);

    event TokenSold(
        address indexed voucherNftContract,
        uint256 indexed tokenId,
        uint256 indexed marketId,
        uint256 amount,
        address buyer,
        address seller,
        uint256 price,
        uint256 totalPrice);

    function verifyVoucherNftContract(address _voucherNftContract) external operatorsOnly {
        voucherNftContractMap[_voucherNftContract] = true;
        emit VerificationVoucherNftContract(_voucherNftContract, true);
    }

    function unVerifyVoucherNftContract(address _voucherNftContract) external operatorsOnly {
        voucherNftContractMap[_voucherNftContract] = false;
        emit VerificationVoucherNftContract(_voucherNftContract, false);
    }

    constructor(
        address _whitelistManager,
        address _operatorManager,
        address _feeManager){
        whitelistManager = _whitelistManager;
        operatorManager = _operatorManager;
        feeManager = _feeManager;
    }

    function place(uint256 _amount, uint256 _tokenId, address _voucherNftContract, uint256 _perTokenPrice) external {
        // check voucher contract verification
        require(voucherNftContractMap[_voucherNftContract], "Not Valid Voucher Contract");
        // increment marketId
        _marketItemIds.increment();
        uint256 marketId = _marketItemIds.current();
        // create market map for each marketId
        _marketItemMap[marketId] = MarketItem(
            _tokenId,
            _voucherNftContract,
            _amount,
            _perTokenPrice,
            address(msg.sender)
        );
        // nft : transfer from seller wallet to contract
        IERC1155(_voucherNftContract).safeTransferFrom(msg.sender, address(this), _tokenId, _amount, "");
        // trigger event
        emit TokenPlaced(_voucherNftContract, _tokenId, marketId, _amount, address(msg.sender), _perTokenPrice);
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
        IERC1155(marketItem.voucherNftContract).safeTransferFrom(address(this), msg.sender, marketItem.tokenId, _amount, "");
        // trigger event
        emit TokenUnPlaced(marketItem.voucherNftContract, marketItem.tokenId, _marketId, _amount, marketItem.amount, marketItem.seller, marketItem.price);
    }

    function transferByOperator(uint256 _marketId, uint256 _amount, address _receiver) external operatorsOnly {
        // check amount must be higher than 0
        require(_amount > 0, "Must be higher than zero");
        // get marketItem
        MarketItem storage marketItem = _marketItemMap[_marketId];
        // check amount from market amount
        require(marketItem.amount >= _amount, "Not Enough amount");
        // divide totalPrice to decimal
        uint256 totalPrice = marketItem.price.mul(_amount).div(10 ** 18);
        // check contract balance
        require(IERC1155(marketItem.voucherNftContract).balanceOf(address(this), marketItem.tokenId) >= _amount, "Insufficient tokens in contract");
        // deduct amount
        marketItem.amount = marketItem.amount.sub(_amount);
        // company fee amount
        uint256 calculatedAmount = IFeeManager(feeManager).mMarketFeeAmount(_amount);
        uint256 remainAmount = _amount.sub(calculatedAmount);
        // transfer voucher token to buyer
        IERC1155(marketItem.voucherNftContract).safeTransferFrom(address(this), _receiver, marketItem.tokenId, remainAmount, "");
        // transfer voucher token to company
        IERC1155(marketItem.voucherNftContract).safeTransferFrom(address(this), IFeeManager(feeManager).feeAddress(), marketItem.tokenId, calculatedAmount, "");
        // trigger event
        emit TokenSold(marketItem.voucherNftContract, marketItem.tokenId, _marketId, _amount, _receiver, marketItem.seller, marketItem.price, totalPrice);
    }

}
