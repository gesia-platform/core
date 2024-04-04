// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable.sol";

contract NetworkAccount is Ownable {
    struct Gateway {
        string hostname;
        uint port; // http, 0 if not supported
        uint sslPort; // https, 0 if not supported
        uint wsPort; // ws, 0 if not supported
        uint wssPort; // wss, 0 if not supported
    }

    string public name;
    string public details;

    Gateway public gateway;

    constructor(string memory _name, string memory _details) {
        name = _name;
        details = _details;
    }

    function set(
        string memory _name,
        string memory _details,
        string memory hostname,
        uint port,
        uint sslPort,
        uint wsPort,
        uint wssPort
    ) public onlyOwner {
        name = _name;
        details = _details;

        gateway = Gateway(hostname, port, sslPort, wsPort, wssPort);
    }

    function get()
        public
        view
        returns (
            string memory,
            string memory,
            string memory hostname,
            uint port,
            uint sslPort,
            uint wsPort,
            uint wssPort
        )
    {
        return (
            name,
            details,
            gateway.hostname,
            gateway.port,
            gateway.sslPort,
            gateway.wsPort,
            gateway.wssPort
        );
    }
}
