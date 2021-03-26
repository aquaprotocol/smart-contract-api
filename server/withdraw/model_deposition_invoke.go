package withdraw

type InvokeWithdrawBody struct {
	InvokeWithdraw InvokeWithdraw `json:"invokeDeposition"`
}
type InvokeWithdraw struct {
	DepositionType string `json:"depType"`
	AccountNumberFrom string `json:"accountNumberFrom"`
	AccountNumberTo string `json:"accountNumberTo"`
	Amount string `json:"amount"`
	EthAccount string `json:"ethAccount"`
}