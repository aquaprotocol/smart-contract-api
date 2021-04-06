package insurance

import (
	db "../../../db"
	"log"
	"time"
)

func ProvideInsuranceWorldStateModification(operator string, iIR *InvestInsuranceRequest) {
	dBase := db.ProvideDBConnection()

	insForm, err := dBase.Prepare("INSERT INTO insurance_state(operator, amount, name, type_of, transaction_date) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	r, err := insForm.Exec(operator, iIR.Price, iIR.Name, iIR.TypeOf, time.Now())

	if err != nil {

		log.Fatal(err)
		log.Fatal(r)
	}

	log.Println("INSERT: Operator: " + operator + " price: " + iIR.Price)


	defer dBase.Close()
}

func ProvideInsuranceWorldState(operator string) []*InvestInsuranceResponse {
	dBase := db.ProvideDBConnection()

	results, err := dBase.Query("SELECT operator, amount, name, type_of, transaction_date FROM insurance_state where operator =?", 	operator)

	insurances := make([]*InvestInsuranceResponse, 0)
	if err != nil {
		log.Fatal(err)
	}


	for results.Next() {
		insurance := new(InvestInsuranceResponse)
		err = results.Scan(&insurance.Operator, &insurance.Price, &insurance.Name, &insurance.TypeOf, &insurance.TransactionDate)

		if err != nil {
			log.Fatal(err)
		}

		insurances = append(insurances,insurance)

	}

	return insurances

}

