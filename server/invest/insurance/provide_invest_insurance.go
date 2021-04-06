package insurance

import (
	con "../../../connector"
	res "../../../server/response"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func ProvideInvestInsurance(iIR *InvestInsuranceRequest) *res.InvokeResponseInfo {

	client := con.ClientFactory()

	transactOpts := con.WalletFactory(client)

	instance := con.ContractFactory(client)

	amount := new(big.Int)
	amount.SetString(iIR.Price, 10)

	//asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, name string, _typeOf string, _date string
	tx, err := instance.BuyInsurance(transactOpts,
		common.HexToAddress("ASSET_ADDRESS"),
		common.HexToAddress("TOKEN_ADDRESS"),
		common.HexToAddress("OPERATOR_ADDRESS"),
		amount,
		iIR.Name,
		iIR.TypeOf,
		iIR.Date)

	if err != nil {
		log.Fatal(err)
	}

	ProvideInsuranceWorldStateModification("OPERATOR_ADDRESS", iIR)
	return &res.InvokeResponseInfo{Hash: tx.Hash().Hex()}

}


func GetInsuranceWorldState() []*InvestInsuranceResponse {

	return ProvideInsuranceWorldState("OPERATOR_ADDRESS")
}