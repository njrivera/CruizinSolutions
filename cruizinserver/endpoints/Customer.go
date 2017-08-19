package endpoints

import (
	"encoding/json"
	"log"
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
	customers, err := dbcontext.GetCustomers()
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(customers, w)
}

func createCustomerHandler(r *http.Request, w http.ResponseWriter) {
	customer := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Println(err)
		util.JSONEncode("Unable to add customer", w)
		return
	}
	err = dbcontext.CreateCustomer(customer)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(customer, w)
}

func getCustomerHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	customer, err := dbcontext.GetCustomer(cid)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(customer, w)
}

func deleteCustomerHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	err := dbcontext.DeleteCustomer(cid)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
}

func updateCustomerHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	cid, _ := strconv.Atoi(params["cid"])
	customer := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Println(err)
		util.JSONEncode("Unable to update customer", w)
		return
	}
	customer.Cid = cid
	err = dbcontext.UpdateCustomer(customer)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(customer, w)
}
