package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CruizinSolutions/cruizinserver/models"

	"github.com/CruizinSolutions/cruizinserver/dbcontext"
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/go-martini/martini"
)

func RegisterOrderEndpoints(m *martini.ClassicMartini) {
	m.Group("/orders", func(r martini.Router) {
		r.Post("", createOrderHandler)
		r.Get("/:cid", getCustVehiclesHandler)
	})
}

func createOrderHandler(r *http.Request, w http.ResponseWriter) {
	orderwithitems := models.OrderWithItems{}
	err := json.NewDecoder(r.Body).Decode(&orderwithitems)
	util.CheckErr(err)
	dbcontext.CreateOrder(orderwithitems.Order, orderwithitems.Items)
	util.JSONEncode(orderwithitems, w)
}

func getCustVehiclesHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	vehicles := dbcontext.GetCustVehicles(cid)
	util.JSONEncode(vehicles, w)
}
