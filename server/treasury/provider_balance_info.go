package treasury

import (
	token "../contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func ProvideBalanceOfInfo(info AccountInfo) BalanceInfo {

	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("")
	instance, err := token.NewContract(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(info.AccountNumber)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err, ". Cannot fetch balance from erc20 instance for contractAddress ")

	}

	fbal := new(big.Float)


	fbal.SetString(bal.String())

	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(18)))

	return BalanceInfo{Balance: value.String()}

}
