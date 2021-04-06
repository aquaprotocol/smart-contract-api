package art

import (
	con "../../../connector"
	res "../../../server/response"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func ProvideInvestArt(inArtReq *InvestArtRequest) *res.InvokeResponseInfo {

	client := con.ClientFactory()

	transactOpts := con.WalletFactory(client)

	instance := con.ContractFactory(client)

	amount := new(big.Int)
	amount.SetString(inArtReq.Price, 10)

	//asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, name string
	tx, err := instance.BuyArt(transactOpts,
		common.HexToAddress("ASSET_ADDRESS"),
		common.HexToAddress("TOKEN_ADDRESS"),
		common.HexToAddress("OPERATOR_ADDRESS"),
		amount,
		inArtReq.Name)

	if err != nil {
		log.Fatal(err)
	}

	ProvideArtWorldStateModification("OPERATOR_ADDRESS",inArtReq)

	return &res.InvokeResponseInfo{Hash: tx.Hash().Hex()}

	// tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
}

func GetArtWorldState() []*InvestArtResponse {

	return ProvideArtWorldState("OPERATOR_ADDRESS")
}
