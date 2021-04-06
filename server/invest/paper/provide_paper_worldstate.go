package paper

import (
	db "../../../db"
	"log"
	"time"
)

func ProvidePaperWorldStateModification(operator string, iPR *InvestPaperRequest) {
	dBase := db.ProvideDBConnection()



	insForm, err := dBase.Prepare("INSERT INTO paper_state(operator, issuer_id, amount, series, number_from, number_to, transaction_date) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	r, err := insForm.Exec(operator, iPR.IssuerId, iPR.Price, iPR.Series, iPR.NumberFrom, iPR.NumberTo, time.Now())

	if err != nil {

		log.Fatal(err)
		log.Fatal(r)
	}

	log.Println("INSERT: Operator: " + operator + " price: " + iPR.Price)


	defer dBase.Close()
}

func GetPaperWorldState() []*InvestPaperResponse {

	return ProvidePaperWorldState("OPERATOR_ADDRESS")
}

