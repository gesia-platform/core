// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IPrice {
    function getCarbonPrice(uint256 tokenId) external view returns (uint256);
}
