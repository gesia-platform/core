// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IFeeManager {
    function feeAmount(uint256 _amount) external view returns (uint256);
    function feeAddress() external view returns (address);
}
