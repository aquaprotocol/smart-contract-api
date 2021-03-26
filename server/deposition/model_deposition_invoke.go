package deposition

type InvokeDepositionBody struct {
	InvokeDeposition InvokeDeposition `json:"invokeDeposition"`
}
type InvokeDeposition struct {
	DepositionType string `json:"depType"`
	AccountNumberFrom string `json:"accountNumberFrom"`
	AccountNumberTo string `json:"accountNumberTo"`
	Amount string `json:"amount"`
	EthAccount string `json:"ethAccount"`
}