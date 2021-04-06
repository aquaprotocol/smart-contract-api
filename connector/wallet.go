package connector

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
)

func WalletFactory(client *ethclient.Client) *bind.TransactOpts {

	file := "./tmp/WALLET_FILE"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	password := "PASSWORD"
	account, err := ks.Import(jsonBytes, password, password)

	nonce, err := client.PendingNonceAt(context.Background(), account.Address)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	ks.Unlock(account,"PASSWORD")

	fmt.Println(account.Address.Hex())

	transactOpts, err := bind.NewKeyStoreTransactor(ks,account)

	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = big.NewInt(0)     // in wei
	transactOpts.GasLimit = uint64(797400) // in units
	transactOpts.GasPrice = gasPrice

	return transactOpts
}
