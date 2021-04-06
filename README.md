# smart-contract-api
The Connector between webpage and blockchain's world.
Web service is writen in GO language. 
If it is your first step with GO please read this documentation https://goethereumbook.org/en/



To run `go build`


### Base function:
In the bottom is a short description of the smart-contract-api function is described.
#### GetBalanceOf
The function for obtaining bank account balance.

Request:
```go
type AccountInfo struct {
	AccountNumber string `json:"accountNumber"`
}
```
Response:
```go
type BalanceInfo struct {
	Balance string `json:"balance"`
}
```

#### InvokeDeposition
The function for invoke fiat/crypto currency deposition.

Request:
```go
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
```
Response:
```go
type InvokeResponseInfo struct {
	Hash string `json:"hash"`
}
```
#### InvokeWithdraw
The function for invoke fiat/crypto currency withdraw.

Request:
```go
type InvokeWithdrawBody struct {
	InvokeWithdraw InvokeWithdraw `json:"invokeWithdraw"`
}
type InvokeWithdraw struct {
	DepositionType string `json:"depType"`
	AccountNumberFrom string `json:"accountNumberFrom"`
	AccountNumberTo string `json:"accountNumberTo"`
	Amount string `json:"amount"`
	EthAccount string `json:"ethAccount"`
}
```
Response:
```go
type InvokeResponseInfo struct {
	Hash string `json:"hash"`
}
```

#### InvestArt
The function for buy Art.

Request:
```go
type InvestArtRequest struct {
	Name string `json:"name"`
	Price string `json:"price"`
}
```
Response:
```go
type InvokeResponseInfo struct {
	Hash string `json:"hash"`
}
```

#### GetArtWorldState
The function for get Art world state.

Request:

Response:
```go
type InvestArtResponse struct {
	Operator string `json:"operator"`
	Name string `json:"name"`
	Price string `json:"amount"`
	TransactionDate string `json:"transaction_date"`
}
```


#### InvestInsurance
The function for buy insurance.

Request:
```go
type InvestInsuranceRequest struct {
	Name string `json:"name"`
	Price string `json:"price"`
	TypeOf string `json:"typeOf"`
	Date string `json:"date"`
}
```

Response:
```go
type InvokeResponseInfo struct {
	Hash string `json:"hash"`
}
```

#### GetInsuranceWorldState
The function for get insurance world state.

Request:

Response:
```go
type InvestInsuranceResponse struct {
	Operator string `json:"operator"`
	Price string `json:"amount"`
	Name string `json:"name"`
	TypeOf string `json:"typeOf"`
	TransactionDate string `json:"transaction_date"`
}
```

#### InvestPaper
The function for buy paper.

Request:
```go
type InvestPaperRequest struct {
	IssuerId string `json:"issuerId"`
	Price string `json:"price"`
	Series string `json:"series"`
	NumberFrom string `json:"numberFrom"`
	NumberTo string `json:"numberTo"`
}
```

Response:
```go
type InvokeResponseInfo struct {
	Hash string `json:"hash"`
}
```

#### GetPaperWorldState
The function for get insurance world state.

Request:

Response:
```go
type InvestPaperResponse struct {
	Operator string `json:"operator"`
	IssuerId string `json:"issuerId"`
	Price string `json:"amount"`
	Series string `json:"series"`
	NumberFrom string `json:"numberFrom"`
	NumberTo string `json:"numberTo"`
	TransactionDate string `json:"transaction_date"`
}
```












