// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

import "./AppPermissionBase.sol";
import "../network/NetworkAccount.sol";

contract AppPermissionNetworkAccess is AppPermissionBase {
    struct NetworkAccessResponse {
        bytes signature;
        bool isGranted;
        address notaryAccount;
    }

    mapping(uint256 appID => mapping(address networkAccount => string ip))
        internal networkAccessRequests;

    mapping(string ip => uint256 appID) internal getAppIDByIP;

    mapping(uint256 appID => mapping(address networkAccount => NetworkAccessResponse))
        internal networkAccessResponses;

    event NetworkAccessPermissionRequested(
        uint256 appID,
        address networkAccount,
        string ip
    );

    event NetworkAccessPermissionResponsed(
        uint256 appID,
        address networkAccount,
        bytes signature,
        bool isGranted,
        address notaryAccount
    );

    function requestNetworkAccess(
        uint256 appID,
        address networkAccount,
        string memory ip
    ) external {
        require(
            bytes(networkAccessRequests[appID][networkAccount]).length == 0,
            "already permission requested"
        );

        (uint256 id, , address owner) = appStore.getApp(appID);
        require(owner == msg.sender, "sender is not app owner");

        networkAccessRequests[appID][networkAccount] = ip;
        getAppIDByIP[ip] = appID;

        emit NetworkAccessPermissionRequested(id, networkAccount, ip);
    }

    function setNetworkAccessRequestIP(
        uint256 appID,
        address networkAccount,
        string memory ip
    ) external {
        string memory previousIP = networkAccessRequests[appID][networkAccount];
        require(bytes(previousIP).length != 0, "permission not requested");

        (, , address owner) = appStore.getApp(appID);
        require(owner == msg.sender, "sender is not app owner");

        delete getAppIDByIP[previousIP];

        networkAccessRequests[appID][networkAccount] = ip;
        getAppIDByIP[ip] = appID;
    }

    function responseNetworkAccess(
        uint256 appID,
        address _networkAccount,
        bytes memory signature,
        bool isGranted,
        address notaryAccount
    ) external {
        NetworkAccount networkAccount = NetworkAccount(_networkAccount);
        require(
            networkAccount.owner() == msg.sender,
            "sender is not network account owner"
        );

        networkAccessResponses[appID][_networkAccount] = NetworkAccessResponse(
            signature,
            isGranted,
            notaryAccount
        );

        emit NetworkAccessPermissionResponsed(
            appID,
            _networkAccount,
            signature,
            isGranted,
            notaryAccount
        );
    }

    function getNetworkAccessRequest(
        uint256 appID,
        address networkAccount
    ) public view returns (bool, string memory) {
        string memory ip = networkAccessRequests[appID][networkAccount];

        return (bytes(ip).length != 0, ip);
    }

    function getAppID(string memory ip) public view returns (bool, uint256) {
        uint256 appID = getAppIDByIP[ip];

        return (appID != 0, appID);
    }

    function getNetworkAccessResponse(
        uint256 appID,
        address networkAccount
    ) public view returns (bool responsed, bytes memory, bool) {
        NetworkAccessResponse memory response = networkAccessResponses[appID][
            networkAccount
        ];

        return (
            bytes(response.signature).length != 0,
            response.signature,
            response.isGranted
        );
    }
}
