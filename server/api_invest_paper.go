package server


import (
	invest "./invest/paper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InvestPaper(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("InvestArt request body error")
	}

	var iA invest.InvestPaperRequest
	err = json.Unmarshal(b, &iA)

	fmt.Println("InvestPaper name ", iA.IssuerId, "price", iA.Price, "series", iA.Series, "paperFrom", iA.NumberFrom, "paperTo", iA.NumberTo)

	pIA := invest.ProvideInvestPaper(&iA)

	jIA, _:= json.Marshal(pIA)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jIA)
}

