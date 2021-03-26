package deposition

import (
	con "../../connector"
	res "../../server/response"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

func ProvideInvokeDeposition(invokeDeposition *InvokeDeposition) *res.InvokeResponseInfo {

	client := con.ClientFactory()

	transactOpts := con.WalletFactory(client)

	instance := con.ContractFactory(client)

	amount := new(big.Int)
	assetAddress := common.HexToAddress("")
	amount.SetString(invokeDeposition.Amount, 10)

	var tx *types.Transaction
	var err error
	if invokeDeposition.DepositionType == "C" {
		fmt.Println("Deposit for crypto")
		//n.SendNotification("Crypto deposition")
tx, err = instance.Deposit(transactOpts, assetAddress, amount, common.HexToAddress(invokeDeposition.AccountNumberFrom))

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Deposit for fiat")
		//n.SendNotification("Fiat deposition")
		tx, _ = instance.DepositFiat(transactOpts, assetAddress, amount, invokeDeposition.AccountNumberFrom, common.HexToAddress(invokeDeposition.EthAccount))
	}

	if tx == nil {

	}

	fmt.Println(tx.Hash().Hex())

	return &res.InvokeResponseInfo{Hash: tx.Hash().Hex()}

}
