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
		r.Get("/:cid", getOrdersHandler)
		r.Get("/items/:ordernum", getItemOrdersHandler)
		r.Get("/vehicles/:cid", getCustVehiclesHandler)
	})
}

func createOrderHandler(r *http.Request, w http.ResponseWriter) {
	orderwithitems := models.OrderWithItems{}
	err := json.NewDecoder(r.Body).Decode(&orderwithitems)
	util.CheckErr(err)
	orderwithitems.Order.OrderNum = dbcontext.CreateOrder(orderwithitems.Order, orderwithitems.Items)
	for _, item := range orderwithitems.Items {
		dbcontext.UpdateTireQty(item.ItemNum, item.Qty*-1)
		dbcontext.UpdateRimQty(item.ItemNum, item.Qty*-1)
	}
	util.JSONEncode(orderwithitems, w)
}

func getOrdersHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	orders := dbcontext.GetOrders(cid)
	util.JSONEncode(orders, w)
}

func getItemOrdersHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	ordernum, _ := strconv.Atoi(params["ordernum"])
	itemorders := dbcontext.GetItemOrders(ordernum)
	util.JSONEncode(itemorders, w)
}

func getCustVehiclesHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	vehicles := dbcontext.GetCustVehicles(cid)
	util.JSONEncode(vehicles, w)
}
