// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "../openzeppelin//access/Ownable.sol";
import "./IFeeManager.sol";
import "../operator/IOperator.sol";
import "../openzeppelin//utils/math/SafeMath.sol";

contract FeeManager is IFeeManager {
    using SafeMath for uint256;
    // fee percentages 5 -> 0.5% , 50 -> 5%, 500 -> 50%
    uint256 public voucherMintFeePer;
    uint256 public voucherTransferFeePer;
    uint256 public carbonTransferFeePer;
    uint256 public carbonMintFeePer;
    uint256 public carbonGiftFeePer;
    uint256 public mMarketFeePer;
    uint256 public voucherMarketFeePer;
    uint256 public nftMarketFeePer;
    address public operatorManager;
    address public feeWalletAddress;

    event ChangeFeeAddress(address _feeAddress);
    event ChangeFeePercentage(uint256 _voucherMintFeePer,
        uint256 _voucherTransferPer,
        uint256 _carbonTransferFeePer,
        uint256 _carbonMintFeePer,
        uint256 _carbonGiftFeePer,
        uint256 _mMarketFeePer,
        uint256 _voucherMarketFeePer,
        uint256 _nftMarketFeePer);

    modifier operatorsOnly() {
        require(IOperator(operatorManager).isOperator(msg.sender), "#operatorsOnly:");
        _;
    }

    constructor(
        address _operatorManager,
        address _feeAddress,
        uint256 _voucherMintFeePer,
        uint256 _voucherTransferPer,
        uint256 _carbonTransferFeePer,
        uint256 _carbonMintFeePer,
        uint256 _carbonGiftFeePer,
        uint256 _mMarketFeePer,
        uint256 _voucherMarketFeePer,
        uint256 _nftMarketFeePer){
        operatorManager = _operatorManager;
        feeWalletAddress = _feeAddress;
        voucherMintFeePer = _voucherMintFeePer;
        voucherTransferFeePer = _voucherTransferPer;
        carbonTransferFeePer = _carbonTransferFeePer;
        carbonMintFeePer = _carbonMintFeePer;
        carbonGiftFeePer = _carbonGiftFeePer;
        mMarketFeePer = _mMarketFeePer;
        voucherMarketFeePer = _voucherMarketFeePer;
        nftMarketFeePer = _nftMarketFeePer;
    }

    function changeFeeAddress(address _feeAddress) external operatorsOnly {
        feeWalletAddress = _feeAddress;
        emit ChangeFeeAddress(_feeAddress);
    }

    function changeFeePercentage(
        uint256 _voucherMintFeePer,
        uint256 _voucherTransferPer,
        uint256 _carbonTransferFeePer,
        uint256 _carbonMintFeePer,
        uint256 _carbonGiftFeePer,
        uint256 _mMarketFeePer,
        uint256 _voucherMarketFeePer,
        uint256 _nftMarketFeePer) external operatorsOnly {
        voucherMintFeePer = _voucherMintFeePer;
        voucherTransferFeePer = _voucherTransferPer;
        carbonTransferFeePer = _carbonTransferFeePer;
        carbonMintFeePer = _carbonMintFeePer;
        carbonGiftFeePer = _carbonGiftFeePer;
        mMarketFeePer = _mMarketFeePer;
        voucherMarketFeePer = _voucherMarketFeePer;
        nftMarketFeePer = _nftMarketFeePer;
        emit ChangeFeePercentage(voucherMintFeePer, voucherTransferFeePer, carbonTransferFeePer, carbonMintFeePer, carbonGiftFeePer, mMarketFeePer, voucherMarketFeePer, nftMarketFeePer);
    }

    function voucherMintFeeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(voucherMintFeePer).div(1000);
    }

    function voucherTransferFeeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(voucherTransferFeePer).div(1000);
    }

    function carbonMintFeeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(carbonMintFeePer).div(1000);
    }

    function carbonTransferFeeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(carbonTransferFeePer).div(1000);
    }

    function carbonGiftFeeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(carbonGiftFeePer).div(1000);
    }

    function mMarketFeeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(mMarketFeePer).div(1000);
    }

    function voucherMarketFeeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(voucherMarketFeePer).div(1000);
    }

    function nftMarketFeeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(nftMarketFeePer).div(1000);
    }


    function feeAddress() external view override returns (address) {
        return feeWalletAddress;
    }

}
