// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

import "./AppPermissionBase.sol";
import "../network/NetworkAccount.sol";

contract AppPermissionNetworkAccess is AppPermissionBase {
    struct NetworkAccessRequest {
        string ip;
    }

    struct NetworkAccessResponse {
        bytes signature;
        bool isGranted;
    }

    mapping(uint256 appID => mapping(address networkAccount => NetworkAccessRequest))
        internal networkAccessRequests;

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
        bool isGranted
    );

    function requestNetworkAccess(
        uint256 appID,
        address networkAccount,
        string memory ip
    ) external {
        require(
            bytes(networkAccessRequests[appID][networkAccount].ip).length == 0,
            "already permission requested"
        );

        (uint256 id, , address owner) = appStore.getApp(appID);
        require(owner == msg.sender, "sender is not app owner");

        networkAccessRequests[appID][networkAccount] = NetworkAccessRequest(ip);

        emit NetworkAccessPermissionRequested(id, networkAccount, ip);
    }

    function setNetworkAccessRequestIP(
        uint256 appID,
        address networkAccount,
        string memory ip
    ) external {
        require(
            bytes(networkAccessRequests[appID][networkAccount].ip).length != 0,
            "permission not requested"
        );

        (, , address owner) = appStore.getApp(appID);
        require(owner == msg.sender, "sender is not app owner");

        networkAccessRequests[appID][networkAccount] = NetworkAccessRequest(ip);
    }

    function responseNetworkAccess(
        uint256 appID,
        address _networkAccount,
        bytes memory signature,
        bool isGranted
    ) external {
        NetworkAccount networkAccount = NetworkAccount(_networkAccount);
        require(
            networkAccount.owner() == msg.sender,
            "sender is not network account owner"
        );

        networkAccessResponses[appID][_networkAccount] = NetworkAccessResponse(
            signature,
            isGranted
        );

        emit NetworkAccessPermissionResponsed(
            appID,
            _networkAccount,
            signature,
            isGranted
        );
    }

    function getNetworkAccessRequest(
        uint256 appID,
        address networkAccount
    ) public view returns (bool, string memory) {
        NetworkAccessRequest memory request = networkAccessRequests[appID][
            networkAccount
        ];

        return (bytes(request.ip).length != 0, request.ip);
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
