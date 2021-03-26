package paper

type InvestPaperRequest struct {

	IssuerId string `json:"issuerId"`
	Price string `json:"price"`
	Series string `json:"series"`
	NumberFrom string `json:"numberFrom"`
	NumberTo string `json:"numberTo"`

}
