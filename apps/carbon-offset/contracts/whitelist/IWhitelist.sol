// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

interface IWhitelist {
    function isWhitelist(address _voucherContact,uint256 _tokenId, address _account) external view returns (bool);
}
