// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "../../notary/NotaryPublic.sol";
import "./CarbonEmissionsVoucher.sol";

contract CarbonEmissionsFactory {
    address public immutable notaryPublic;

    event VoucherCreated(address carbonEmissionsVoucher);

    constructor(address _notaryPublic) {
        notaryPublic = _notaryPublic;
    }

    function createVoucher(
        string memory voucherName
    ) external returns (address voucher) {
        require(
            NotaryPublic(notaryPublic).hasMembership(msg.sender),
            "sender is not notary public member"
        );

        voucher = address(
            new CarbonEmissionsVoucher(voucherName, address(this))
        );

        emit VoucherCreated(voucher);

        return voucher;
    }
}
