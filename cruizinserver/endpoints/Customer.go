package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/CruizinSolutions/cruizinserver/dbcontext"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/go-martini/martini"
)

func RegisterCustomerEndpoints(m *martini.ClassicMartini) {
	m.Group("/customers", func(r martini.Router) {
		r.Get("", getCustomersHandler)
		r.Post("", createCustomerHandler)
		r.Get("/:cid", getCustomerHandler)
		r.Delete("/:cid", deleteCustomerHandler)
		r.Put("/:cid", updateCustomerHandler)
	})
}

func getCustomersHandler(r *http.Request, w http.ResponseWriter) {
	customers := dbcontext.GetCustomers()
	util.JSONEncode(customers, w)
}

func createCustomerHandler(r *http.Request, w http.ResponseWriter) {
	customer := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	util.CheckErr(err)
	dbcontext.CreateCustomer(customer)
	util.JSONEncode(customer, w)
}

func getCustomerHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	customer := dbcontext.GetCustomer(cid)
	util.JSONEncode(customer, w)
}

func deleteCustomerHandler(r *http.Request, params martini.Params) {
	cid, _ := strconv.Atoi(params["cid"])
	dbcontext.DeleteCustomer(cid)
}

func updateCustomerHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	customer := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	util.CheckErr(err)
	customer.Cid = cid
	dbcontext.UpdateCustomer(customer)
	util.JSONEncode(customer, w)
}
