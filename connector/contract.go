package connector

import (
	token "../server/aqua"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

const contractAddress = ""

func ContractFactory(client *ethclient.Client) *token.Aqua {
	contractAddress := common.HexToAddress(contractAddress)

	instance, err := token.NewAqua(contractAddress, client)

	if err != nil {
		log.Fatal(err)
	}
	return instance

}
