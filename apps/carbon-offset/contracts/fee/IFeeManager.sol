// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IFeeManager {
    function voucherMintFeeAmount(uint256 _amount) external view returns (uint256);
    function voucherTransferFeeAmount(uint256 _amount) external view returns (uint256);
    function carbonMintFeeAmount(uint256 _amount) external view returns (uint256);
    function carbonTransferFeeAmount(uint256 _amount) external view returns (uint256);
    function carbonGiftFeeAmount(uint256 _amount) external view returns (uint256);
    function mMarketFeeAmount(uint256 _amount) external view returns (uint256);
    function voucherMarketFeeAmount(uint256 _amount) external view returns (uint256);
    function nftMarketFeeAmount(uint256 _amount) external view returns (uint256);
    function feeAddress() external view returns (address);
}
