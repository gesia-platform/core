// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "./NotaryAccount.sol";

abstract contract NotaryKeeper {
    mapping(address => bool) public authorizations;

    address public immutable notaryPublic;

    modifier onlyNotaryPublic() {
        require(msg.sender == notaryPublic, "sender is not notary public.");
        _;
    }

    modifier onlyAuthorized() {
        require(
            authorizations[msg.sender] == true,
            "sender is not authorized."
        );
        _;
    }

    constructor(address _notaryPublic) {
        notaryPublic = _notaryPublic;
    }

    function authorize(address account) external onlyNotaryPublic {
        authorizations[account] = true;
    }

    function revoke(address account) external onlyNotaryPublic {
        delete authorizations[account];
    }

    function getNotaryPublic() external view returns (address) {
        return notaryPublic;
    }
}
