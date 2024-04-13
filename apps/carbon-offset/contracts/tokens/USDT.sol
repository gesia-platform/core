// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../openzeppelin//token/ERC20/ERC20.sol";

contract USDT is ERC20 {
    constructor() ERC20("Tether", "USDT") {
        super._mint(_msgSender(), 30000000000 ether);
    }

    function decimals() public view virtual override returns (uint8) {
        return 6;
    }
}
