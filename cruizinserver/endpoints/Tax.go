package endpoints

import (
	"os"
	"github.com/CruizinSolutions/cruizinserver/util"
	"net/http"
	"github.com/go-martini/martini"
)

func RegisterTaxEndpoint(m *martini.ClassicMartini) {
	m.Get("/taxRate", getTaxRate)
}

func getTaxRate(r *http.Request, w http.ResponseWriter) {
	util.JSONEncode(os.Getenv("TAX_RATE"), w)
}