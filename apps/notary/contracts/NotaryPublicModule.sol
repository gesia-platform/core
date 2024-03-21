// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.23;

import "./NotaryPublicDAO.sol";
import "./NotaryAccount.sol";
import "./NotaryModule.sol";

contract NotaryPublicModule is NotaryPublicDAO {
    mapping(address module => address auditor) public moduleAuditors;

    mapping(address module => bool passed) public moduleAuditResults;

    event ModuleAuditRequested(address module, address auditor);

    event ModuleAudited(address module);

    function requestModuleAudit(address module, address auditor) external {
        require(
            moduleAuditors[module] == address(0),
            "module has already been audit requested."
        );

        require(
            NotaryModule(module).notaryPublicAddress() == address(this),
            "invalid public address of module."
        );

        require(getMembership(auditor) == true, "auditor is not DAO member.");

        moduleAuditors[module] = auditor;

        emit ModuleAuditRequested(module, auditor);
    }

    function auditModule(address module) external onlyMember {
        require(
            moduleAuditors[module] != msg.sender,
            "module is not auditable."
        );

        moduleAuditResults[module] = true;

        emit ModuleAudited(module);
    }
}
