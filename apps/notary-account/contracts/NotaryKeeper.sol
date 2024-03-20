// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublic.sol";

contract NotaryKeeper {
    NotaryPublic immutable notaryPublic;

    modifier onlyCallAuthorized(address account) {
        require(
            notaryPublic.keeperCallAuthorized(
                address(this),
                account,
                msg.sender
            ) == true,
            "sender is not call authorized."
        );
        _;
    }

    constructor(address _notaryPublic) {
        notaryPublic = NotaryPublic(_notaryPublic);
    }

    function getNotaryPublic() external view returns (address) {
        return address(notaryPublic);
    }
}
