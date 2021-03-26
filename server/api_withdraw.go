package server

import (
	withdraw "./withdraw"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InvokeWithdraw(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("GetAccount request body error")
	}

	var iD withdraw.InvokeWithdrawBody
	err = json.Unmarshal(b, &iD)

	fmt.Println("InvokeWithdraw account number from request ", iD.InvokeWithdraw.AccountNumberFrom)

	pID := withdraw.ProvideInvokeWithdraw(&iD.InvokeWithdraw)

	jPID, _:= json.Marshal(pID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jPID)
}
