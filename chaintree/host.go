package chaintree

import "github.com/ethereum/go-ethereum/ethclient"

type Host struct {
	client   *ethclient.Client
	wsClient *ethclient.Client
}

func NewHost(url string, wsURL string) Host {
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	wsClient, err := ethclient.Dial(wsURL)
	if err != nil {
		panic(err)
	}

	return Host{client, wsClient}
}

func (host *Host) Client() *ethclient.Client {
	return host.client
}

func (host *Host) WSClient() *ethclient.Client {
	return host.wsClient
}
