package art

type InvestArtRequest struct {
	Name string `json:"name"`
	Price string `json:"price"`
}


type InvestArtResponse struct {
	Operator string `json:"operator"`
	Name string `json:"name"`
	Price string `json:"amount"`
	TransactionDate string `json:"transaction_date"`
}