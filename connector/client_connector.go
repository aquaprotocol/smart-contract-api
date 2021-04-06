package connector

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)


const url = "INFURA_WSS_ADDRESS"

func ClientFactory() *ethclient.Client  {
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
