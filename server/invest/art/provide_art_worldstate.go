package art

import (
	db "../../../db"
	"log"
	"time"
)

func ProvideArtWorldStateModification(operator string, inArtReq *InvestArtRequest) {
	dBase := db.ProvideDBConnection()

	insForm, err := dBase.Prepare("INSERT INTO art_state(operator, amount, name, transaction_date) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	r, err := insForm.Exec(operator, inArtReq.Price, inArtReq.Name, time.Now())

	if err != nil {

		log.Fatal(err)
		log.Fatal(r)
	}

	log.Println("INSERT: Operator: " + operator + " price: " + inArtReq.Price)


	defer dBase.Close()
}

func ProvideArtWorldState(operator string) []*InvestArtResponse {
	dBase := db.ProvideDBConnection()


	results, err := dBase.Query("SELECT operator, amount, name, transaction_date FROM art_state where operator =?", 	operator)

	arts := make([]*InvestArtResponse, 0)
	if err != nil {
		log.Fatal(err)
	}


	for results.Next() {
		art := new(InvestArtResponse)
		err = results.Scan(&art.Operator, &art.Price, &art.Name, &art.TransactionDate)

		if err != nil {
			log.Fatal(err)
		}

		arts = append(arts,art)

	}

	return arts

}

