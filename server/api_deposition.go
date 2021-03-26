package server

import (
	deposition "./deposition"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InvokeDeposition(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("GetAccount request body error")
	}

	var iD deposition.InvokeDepositionBody
	err = json.Unmarshal(b, &iD)

	fmt.Println("InvokeDeposition account number from request ", iD.InvokeDeposition.AccountNumberFrom)

	pID := deposition.ProvideInvokeDeposition(&iD.InvokeDeposition)

	jPID, _:= json.Marshal(pID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jPID)
}
