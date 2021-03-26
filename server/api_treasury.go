package server

import (
	treasury "./treasury"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetBalanceOf(w http.ResponseWriter, r *http.Request) {

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println("GetAccount request body error")
	}

	var aI treasury.AccountInfo
	err = json.Unmarshal(b, &aI)
	fmt.Println("GetAccount account number from request " + aI.AccountNumber)


	balanceInfo := treasury.ProvideBalanceOfInfo(aI)

 	js,_ :=json.Marshal(balanceInfo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}