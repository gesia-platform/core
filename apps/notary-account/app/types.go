package app

import "github.com/ethereum/go-ethereum/common"

type (
	EventNotaryAccountCreated struct {
		Chain         string
		Domain        string
		NotaryAccount common.Address
	}

	EventRPCNotarized struct {
		Notary common.Address
		Url    string
	}

	RPCRequest struct {
		Result string
	}

	RPCResponse struct {
		Result string
	}
)
