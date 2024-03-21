// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicDAO.sol";
import "./NotaryAccount.sol";
import "./NotaryModule.sol";

contract NotaryPublicModule is NotaryPublicDAO {
    mapping(address module => address account) public moduleAuditRequesters;

    mapping(address module => mapping(address member => bool))
        public moduleAuditResults;

    event ModuleAuditRequested(address module, address account);

    event ModuleAudited(address module, address member, bool passed);

    function requestModuleAudit(address module, address account) external {
        require(
            moduleAuditRequesters[account] == address(0),
            "module has already been audit requested."
        );

        require(
            NotaryModule(module).notaryPublicAddress() == address(this),
            "invalid public address of module."
        );

        require(
            NotaryAccount(account).getOwner() == msg.sender,
            "sender is not notary account owner."
        );

        moduleAuditRequesters[module] = account;

        emit ModuleAuditRequested(module, account);
    }

    function auditModule(address module, bool passed) external onlyMember {
        require(
            moduleAuditRequesters[module] != address(0),
            "module is not auditable."
        );

        moduleAuditResults[module][msg.sender] = passed;

        emit ModuleAudited(module, msg.sender, passed);
    }
}
