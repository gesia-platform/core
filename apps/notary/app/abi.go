package app

const (
	NotaryPublicABI = `[
		{
		  "inputs": [],
		  "stateMutability": "nonpayable",
		  "type": "constructor"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "string",
			  "name": "domain",
			  "type": "string"
			},
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "caller",
			  "type": "address"
			},
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "notaryAccount",
			  "type": "address"
			}
		  ],
		  "name": "NotaryAccountUpdated",
		  "type": "event"
		},
		{
		  "inputs": [
			{
			  "internalType": "string",
			  "name": "domain",
			  "type": "string"
			},
			{
			  "internalType": "address",
			  "name": "caller",
			  "type": "address"
			},
			{
			  "internalType": "bytes32[]",
			  "name": "merkleProof",
			  "type": "bytes32[]"
			}
		  ],
		  "name": "createAccount",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getMerkleRoot",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "string",
			  "name": "domain",
			  "type": "string"
			}
		  ],
		  "name": "getNotaryAccount",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "notaryAccount",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getOwner",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "merkleRoot",
		  "outputs": [
			{
			  "internalType": "bytes32",
			  "name": "",
			  "type": "bytes32"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "notaryAccount",
			  "type": "address"
			},
			{
			  "internalType": "bytes",
			  "name": "signature",
			  "type": "bytes"
			},
			{
			  "internalType": "string",
			  "name": "url",
			  "type": "string"
			},
			{
			  "internalType": "bytes32[]",
			  "name": "merkleProof",
			  "type": "bytes32[]"
			}
		  ],
		  "name": "notarizeAccount",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "string",
			  "name": "",
			  "type": "string"
			}
		  ],
		  "name": "notaryAccounts",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "owner",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "notaryAccount",
			  "type": "address"
			},
			{
			  "internalType": "address",
			  "name": "caller",
			  "type": "address"
			},
			{
			  "internalType": "bytes32[]",
			  "name": "merkleProof",
			  "type": "bytes32[]"
			}
		  ],
		  "name": "setAccountCaller",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes32",
			  "name": "_merkleRoot",
			  "type": "bytes32"
			}
		  ],
		  "name": "setMerkleRoot",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_owner",
			  "type": "address"
			}
		  ],
		  "name": "setOwner",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		}
	  ]`

	NotaryAccountABI = `[
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_notaryPublic",
			  "type": "address"
			},
			{
			  "internalType": "string",
			  "name": "_domain",
			  "type": "string"
			},
			{
			  "internalType": "address",
			  "name": "_caller",
			  "type": "address"
			}
		  ],
		  "stateMutability": "nonpayable",
		  "type": "constructor"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "notary",
			  "type": "address"
			}
		  ],
		  "name": "Canceled",
		  "type": "event"
		},
		{
		  "anonymous": false,
		  "inputs": [
			{
			  "indexed": false,
			  "internalType": "address",
			  "name": "notary",
			  "type": "address"
			}
		  ],
		  "name": "Notarized",
		  "type": "event"
		},
		{
		  "inputs": [],
		  "name": "caller",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "signature",
			  "type": "bytes"
			}
		  ],
		  "name": "cancel",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "domain",
		  "outputs": [
			{
			  "internalType": "string",
			  "name": "",
			  "type": "string"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getCaller",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "getDomain",
		  "outputs": [
			{
			  "internalType": "string",
			  "name": "",
			  "type": "string"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "notary",
			  "type": "address"
			}
		  ],
		  "name": "getNotaryURL",
		  "outputs": [
			{
			  "internalType": "string",
			  "name": "",
			  "type": "string"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "bytes",
			  "name": "signature",
			  "type": "bytes"
			},
			{
			  "internalType": "string",
			  "name": "url",
			  "type": "string"
			}
		  ],
		  "name": "notarize",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		},
		{
		  "inputs": [],
		  "name": "notaryPublic",
		  "outputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "",
			  "type": "address"
			}
		  ],
		  "name": "notaryURLs",
		  "outputs": [
			{
			  "internalType": "string",
			  "name": "",
			  "type": "string"
			}
		  ],
		  "stateMutability": "view",
		  "type": "function"
		},
		{
		  "inputs": [
			{
			  "internalType": "address",
			  "name": "_caller",
			  "type": "address"
			}
		  ],
		  "name": "setCaller",
		  "outputs": [],
		  "stateMutability": "nonpayable",
		  "type": "function"
		}
	  ]`
)
