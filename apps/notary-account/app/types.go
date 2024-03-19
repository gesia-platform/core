package app

import "github.com/ethereum/go-ethereum/common"

type (
	EventNotaryAccountUpdated struct {
		Domain        string
		Caller        common.Address
		NotaryAccount common.Address
	}

	EventNotarized struct {
		Notary common.Address
	}

	RPCRequest struct {
		Result string
	}

	RPCResponse struct {
		Result string
	}
)
