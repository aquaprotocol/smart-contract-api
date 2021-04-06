package server


import (
	invest "./invest/art"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func InvestArt(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("InvestArt request body error")
	}

	var iA invest.InvestArtRequest
	err = json.Unmarshal(b, &iA)

	fmt.Println("InvestArt name ", iA.Name, "price", iA.Price)

	pIA := invest.ProvideInvestArt(&iA)

	jIA, _:= json.Marshal(pIA)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jIA)
}

func GetArtWorldState(w http.ResponseWriter, r *http.Request) {

	pIA := invest.GetArtWorldState()

	jIA, _:= json.Marshal(pIA)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(jIA)
}

