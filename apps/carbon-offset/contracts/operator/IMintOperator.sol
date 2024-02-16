// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IMintOperator {
    function mintByOperator(address _receiver, uint256 _amount, uint256 _tokenId, string memory _metadata) external;
}