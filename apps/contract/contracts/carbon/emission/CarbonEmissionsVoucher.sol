// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";

contract CarbonEmissionsVoucher is ERC1155, ERC1155Holder {
    string private name;

    mapping(address => bool) private _minterApprovals;

    address public immutable carbonEmissionsFactory;

    event CarbonEmitted(uint256 applicationID, uint256 emissions, bytes userID);

    /**
     * @param _name The emissions categroy name.
     */

    constructor(
        string memory _name,
        address _carbonEmissionsFactory
    ) ERC1155("") {
        name = _name;
        carbonEmissionsFactory = _carbonEmissionsFactory;
    }

    function mint(
        uint256 applicationID,
        uint256 emissions, // scaled by 10,000
        bytes memory userID
    ) external {
        require(_minterApprovals[msg.sender], "sender is not approved minter");

        super._mint(address(this), applicationID, emissions, "");

        emit CarbonEmitted(applicationID, emissions, userID);
    }

    function setMinterApproval(address minter, bool approval) external {
        require(
            carbonEmissionsFactory == msg.sender,
            "sender is not carbon voucher factory"
        );

        _minterApprovals[minter] = approval;
    }

    function isApprovedMinter(address calculator) public view returns (bool) {
        return _minterApprovals[calculator];
    }

    function getName() public view returns (string memory) {
        return name;
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view virtual override(ERC1155, ERC1155Receiver) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
