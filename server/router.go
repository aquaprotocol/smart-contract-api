package server

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = Routes{

	Route{
		"GetBalanceOf",
		strings.ToUpper("Post"),
		"/treasury/balanceOf",
		GetBalanceOf,
	},

	Route{
		"InvokeDeposition",
		strings.ToUpper("Post"),
		"/deposition/invoke",
		InvokeDeposition,
	},
	Route{
		"InvokeWithdraw",
		strings.ToUpper("Post"),
		"/withdraw/invoke",
		InvokeWithdraw,
	},

	Route{
		"InvestArt",
		strings.ToUpper("Post"),
		"/invest/art",
		InvestArt,
	},

	Route{
		"GetArtWorldState",
		strings.ToUpper("Post"),
		"/invest/getArtState",
		GetArtWorldState,
	},

	Route{
		"InvestInsurance",
		strings.ToUpper("Post"),
		"/invest/insurance",
		InvestInsurance,
	},

	Route{
		"GetInsuranceWorldState",
		strings.ToUpper("Post"),
		"/invest/getInsuranceState",
		GetInsuranceWorldState,
	},

	Route{
		"InvestPaper",
		strings.ToUpper("Post"),
		"/invest/paper",
		InvestPaper,
	},

	Route{
		"GetPaperWorldState",
		strings.ToUpper("Post"),
		"/invest/getPaperState",
		GetPaperWorldState,
	},
}
