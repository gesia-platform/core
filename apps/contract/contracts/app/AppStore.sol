// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";

contract AppStore is Ownable {
    struct App {
        uint256 id;
        string domain;
        address owner;
    }

    uint256 private nextAppID = 1;

    mapping(uint256 => App) internal apps;

    constructor() {}

    event AppCreated(uint256 appID, string domain, address owner);

    event AppOwnershipTransferred(
        uint256 applicationID,
        address previousOwner,
        address newOwner
    );

    function createApp(string memory domain, address owner) external {
        App memory app = App(nextAppID, domain, owner);
        apps[nextAppID] = app;

        emit AppCreated(app.id, app.domain, app.owner);

        nextAppID++;
    }

    function transferAppOwnership(uint256 appID, address newOwner) external {
        App storage app = apps[appID];

        require(msg.sender == app.owner, "sender is not owner.");

        app.owner = newOwner;

        emit AppOwnershipTransferred(appID, msg.sender, newOwner);
    }

    function getApp(
        uint256 appID
    ) public view returns (uint256, string memory, address) {
        App memory app = apps[appID];

        return (app.id, app.domain, app.owner);
    }
}
