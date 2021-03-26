package server


import (
	invest "./invest/insurance"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InvestInsurance(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("InvestArt request body error")
	}

	var iA invest.InvestInsuranceRequest
	err = json.Unmarshal(b, &iA)

	fmt.Println("InvestInsurance name ", iA.Name, "price", iA.Price, "typeOf", iA.TypeOf, "date", iA.Date)

	pIA := invest.ProvideInvestInsurance(&iA)

	jIA, _:= json.Marshal(pIA)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jIA)
}

