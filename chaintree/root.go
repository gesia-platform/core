package chaintree

import "github.com/ethereum/go-ethereum/ethclient"

type Root struct {
	client *ethclient.Client
}

func NewRoot(url string) Root {
	client, err := ethclient.Dial(url)
	if err != nil {
		panic(err)
	}

	return Root{client}
}

func (root *Root) Client() *ethclient.Client {
	return root.client
}
