package endpoints

import (
	"encoding/json"
	"errors"
	"log"
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
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to create order"), w)
		return
	}
	orderwithitems.Order.OrderNum, err = dbcontext.CreateOrder(orderwithitems.Order, orderwithitems.Items)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	for _, item := range orderwithitems.Items {
		dbcontext.UpdateTireQty(item.ItemNum, item.Qty*-1)
		if err != nil {
			util.JSONEncode(err, w)
			return
		}
		dbcontext.UpdateRimQty(item.ItemNum, item.Qty*-1)
		if err != nil {
			util.JSONEncode(err, w)
			return
		}
	}
	util.JSONEncode(orderwithitems, w)
}

func getOrdersHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	orders, err := dbcontext.GetOrders(cid)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(orders, w)
}

func getItemOrdersHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	ordernum, _ := strconv.Atoi(params["ordernum"])
	itemorders, err := dbcontext.GetItemOrders(ordernum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(itemorders, w)
}

func getCustVehiclesHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	vehicles, err := dbcontext.GetCustVehicles(cid)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(vehicles, w)
}
