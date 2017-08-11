package endpoints

import (
	"net/http"
	"strconv"

	"github.com/CruizinSolutions/cruizinserver/dbcontext"
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/go-martini/martini"
)

func RegisterOrderEndpoints(m *martini.ClassicMartini) {
	m.Group("/orders", func(r martini.Router) {
		r.Get("/:cid", getCustVehiclesHandler)
	})
}

func getCustVehiclesHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	vehicles := dbcontext.GetCustVehicles(cid)
	util.JSONEncode(vehicles, w)
}
