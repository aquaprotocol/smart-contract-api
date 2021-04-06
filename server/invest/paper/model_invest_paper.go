package paper

type InvestPaperRequest struct {

	IssuerId string `json:"issuerId"`
	Price string `json:"price"`
	Series string `json:"series"`
	NumberFrom string `json:"numberFrom"`
	NumberTo string `json:"numberTo"`

}

type InvestPaperResponse struct {

	Operator string `json:"operator"`
	IssuerId string `json:"issuerId"`
	Price string `json:"amount"`
	Series string `json:"series"`
	NumberFrom string `json:"numberFrom"`
	NumberTo string `json:"numberTo"`
	TransactionDate string `json:"transaction_date"`

}
