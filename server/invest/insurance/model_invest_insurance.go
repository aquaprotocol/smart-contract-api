package insurance

type InvestInsuranceRequest struct {
	Name string `json:"name"`
	Price string `json:"price"`
	TypeOf string `json:"typeOf"`
	Date string `json:"date"`
}

type InvestInsuranceResponse struct {
	Operator string `json:"operator"`
	Price string `json:"amount"`
	Name string `json:"name"`
	TypeOf string `json:"typeOf"`
	TransactionDate string `json:"transaction_date"`
}
