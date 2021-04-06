package withdraw

import (
	con "../../connector"
	res "../../server/response"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
)

func ProvideInvokeWithdraw(iW *InvokeWithdraw) *res.InvokeResponseInfo {

	client := con.ClientFactory()

	transactOpts := con.WalletFactory(client)

	instance := con.ContractFactory(client)

	fmt.Println("Amount ",iW.Amount)
	amount := new(big.Int)
	assetAddress := common.HexToAddress("ASSET_ADDRESS")
	amount.SetString(iW.Amount, 10)

	ethAddress := common.HexToAddress(iW.EthAccount)

	var tx *types.Transaction
	var err error
	if iW.DepositionType == "C" {
		fmt.Println("Deposit for crypto")
		//currencyAddress := common.HexToAddress(invokeDeposition.EthAccount)

		//n.SendNotification("Crypto deposition")
		//tx, err = instance.Deposit(transactOpts, assetAddress, currencyAddress, amount, common.HexToAddress(invokeDeposition.AccountNumberFrom))

		//asset common.Address, ethAddress common.Address, amount *big.Int, onBehalfOf common.Address
		tx, err = instance.Withdraw(transactOpts, assetAddress, ethAddress, amount, common.HexToAddress(iW.AccountNumberFrom))

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Withdraw for fiat", amount)
		//n.SendNotification("Fiat deposition")
		tx, _ = instance.WithdrawFiat(transactOpts, assetAddress, ethAddress, iW.AccountNumberFrom, amount, common.HexToAddress(iW.EthAccount))
	}

	if tx == nil {

	}

	fmt.Println(tx.Hash().Hex())

	return &res.InvokeResponseInfo{Hash: tx.Hash().Hex()}
}