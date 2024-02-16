// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./IWhitelist.sol";
import "../operator/IOperator.sol";

contract WhitelistManager is IWhitelist {

    mapping(address => mapping(uint256 => mapping(address => bool))) public whitelistMap;
    address public operatorManager;

    event AddWhitelist(address voucher, uint256 tokenId, address account);
    event RemoveWhitelist(address voucher, uint256 tokenId, address account);

    modifier operatorsOnly() {
        require(IOperator(operatorManager).isOperator(msg.sender), "#operatorsOnly:");
        _;
    }

    constructor(address _operatorManager) {
        operatorManager = _operatorManager;
    }

    function addWhitelist(address _voucherContract, uint256 _tokenId, address _account) external operatorsOnly {
        whitelistMap[_voucherContract][_tokenId][_account] = true;
        emit AddWhitelist(_voucherContract, _tokenId, _account);
    }

    function removeWhitelist(address _voucherContract, uint256 _tokenId, address _account) external operatorsOnly {
        whitelistMap[_voucherContract][_tokenId][_account] = false;
        emit RemoveWhitelist(_voucherContract, _tokenId, _account);
    }

    function isWhitelist(address _voucherContract, uint256 _tokenId, address _account) external view override returns (bool) {
        return whitelistMap[_voucherContract][_tokenId][_account];
    }
}
