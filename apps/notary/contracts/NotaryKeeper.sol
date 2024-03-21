// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryAccount.sol";
import "./NotaryModule.sol";

abstract contract NotaryKeeper {
    address internal immutable notaryModule;
    NotaryAccount internal immutable notaryAccount;

    constructor(NotaryModule _notaryModule, NotaryAccount _notaryAccount) {
        notaryModule = address(_notaryModule);
        notaryAccount = _notaryAccount;
    }

    modifier onlyNotaryAccount() {
        require(
            msg.sender == notaryAccount.getOwner(),
            "sender is not notary account."
        );
        _;
    }

    function doFallback(address impl) internal virtual {
        assembly {
            calldatacopy(0, 0, calldatasize())
            let result := delegatecall(gas(), impl, 0, calldatasize(), 0, 0)
            returndatacopy(0, 0, returndatasize())
            switch result
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return(0, returndatasize())
            }
        }
    }

    function notaryAccountAddress() public view returns (address) {
        return address(notaryAccount);
    }

    function notaryModuleAddress() public view returns (address) {
        return notaryModule;
    }
}
