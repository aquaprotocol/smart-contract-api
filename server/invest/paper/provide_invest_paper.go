package paper

import (
	con "../../../connector"
	db "../../../db"
	res "../../../server/response"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func ProvideInvestPaper(iPR *InvestPaperRequest) *res.InvokeResponseInfo {

	client := con.ClientFactory()

	transactOpts := con.WalletFactory(client)

	instance := con.ContractFactory(client)

	amount := new(big.Int)
	amount.SetString(iPR.Price, 10)

	numberFrom:= new(big.Int)
	amount.SetString(iPR.NumberFrom, 10)

	numberTo := new(big.Int)
	amount.SetString(iPR.NumberTo, 10)

	// asset common.Address, tokenAddress common.Address, operator common.Address, amount *big.Int, issuerId string, series string, numberFrom *big.Int, numberTo *big.Int
	tx, err := instance.BuyPaper(transactOpts,
		common.HexToAddress("ASSET_ADDRESS"),
		common.HexToAddress("TOKEN_ADDRESS"),
		common.HexToAddress("OPERATOR_ADDRESS"),
		amount,
		iPR.IssuerId,
		iPR.Series,
		numberFrom,
		numberTo)

	if err != nil {
		log.Fatal(err)
	}

	ProvidePaperWorldStateModification("OPERATOR_ADDRESS", iPR)
	return &res.InvokeResponseInfo{Hash: tx.Hash().Hex()}

}

func ProvidePaperWorldState(operator string) []*InvestPaperResponse {
	dBase := db.ProvideDBConnection()

	results, err := dBase.Query("SELECT operator, issuer_id, amount, series, number_from, number_to, transaction_date FROM paper_state where operator =?", 	operator)

	papers := make([]*InvestPaperResponse, 0)
	if err != nil {
		log.Fatal(err)
	}


	for results.Next() {
		paper := new(InvestPaperResponse)
		err = results.Scan(&paper.Operator, &paper.IssuerId, &paper.Price, &paper.Series, &paper.NumberFrom, &paper.NumberTo, &paper.TransactionDate)

		if err != nil {
			log.Fatal(err)
		}

		papers = append(papers,paper)

	}

	return papers

}