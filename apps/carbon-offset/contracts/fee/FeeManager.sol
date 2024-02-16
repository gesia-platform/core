// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./IFeeManager.sol";
import "../operator/IOperator.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract FeeManager is IFeeManager {
    using SafeMath for uint256;
    uint256 public feePercentage; // 5 -> 0.5% , 50 -> 5%, 500 -> 50%
    address public operatorManager;
    address public feeWalletAddress;

    event ChangeFeeAddress(address _feeAddress);
    event ChangeFeePercentage(uint256 _feePercentage);

    modifier operatorsOnly() {
        require(IOperator(operatorManager).isOperator(msg.sender), "#operatorsOnly:");
        _;
    }

    constructor(address _operatorManager, address _feeAddress, uint256 _feePercentage){
        operatorManager = _operatorManager;
        feeWalletAddress = _feeAddress;
        feePercentage = _feePercentage;
    }

    function changeFeeAddress(address _feeAddress) external operatorsOnly {
        feeWalletAddress = _feeAddress;
        emit ChangeFeeAddress(_feeAddress);
    }

    function changeFeePercentage(uint256 _feePercentage) external operatorsOnly {
        feePercentage = _feePercentage;
        emit ChangeFeePercentage(_feePercentage);
    }

    function feeAmount(uint256 _feeAmount) external view override returns (uint256) {
        return _feeAmount.mul(feePercentage).div(1000);
    }

    function feeAddress() external view override returns (address) {
        return feeWalletAddress;
    }

}
