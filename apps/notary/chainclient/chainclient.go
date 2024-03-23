package chainclient

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type ChainClient struct {
	Local  *ethclient.Client
	Notary *ethclient.Client
}

func NewChainClient(
	url string,
	notaryURL string,
) *ChainClient {
	localClient, err := ethclient.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("failed to dial local ethclient. %d", err))
	}

	notaryClient, err := ethclient.Dial(notaryURL)
	if err != nil {
		panic(fmt.Sprintf("failed to dial notary ethclient. %d", err))
	}

	return &ChainClient{
		Local:  localClient,
		Notary: notaryClient,
	}

}
