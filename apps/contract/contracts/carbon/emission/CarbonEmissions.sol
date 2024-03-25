// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract CarbonEmissions is ERC1155 {
    string private name;

    error InvalidApprovedCalculator();

    /**
     * @param _name The emissions categroy name.
     */

    constructor(string memory _name) ERC1155("") {
        name = _name;
    }

    function mint(
        address to,
        uint256 applicationID,
        uint256 emossions, // scaled by 10,000
        bytes memory userID
    ) external {
        super._mint(to, applicationID, emossions, userID);
    }

    function getName() public view returns (string memory) {
        return name;
    }
}
