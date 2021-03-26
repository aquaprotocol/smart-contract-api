package insurance

type InvestInsuranceRequest struct {
	Name string `json:"name"`
	Price string `json:"price"`
	TypeOf string `json:"typeOf"`
	Date string `json:"date"`
}
