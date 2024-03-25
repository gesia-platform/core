package chaintree

import "github.com/ethereum/go-ethereum/ethclient"

type Host struct {
	client *ethclient.Client
}

func NewHost(url string) Host {
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	return Host{client}
}

func (host *Host) Client() *ethclient.Client {
	return host.client
}
