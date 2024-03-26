package chaintree

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type Root struct {
	client   *ethclient.Client
	wsClient *ethclient.Client
}

func NewRoot(url string, wsURL string) Root {
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	wsClient, err := ethclient.Dial(wsURL)
	if err != nil {
		panic(err)
	}

	return Root{client, wsClient}
}

func (root *Root) Client() *ethclient.Client {
	return root.client
}

func (root *Root) WSClient() *ethclient.Client {
	return root.wsClient
}
