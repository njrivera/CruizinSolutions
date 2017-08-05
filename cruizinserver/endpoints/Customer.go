package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/CruizinSolutions/cruizinserver/dbcontext"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/go-martini/martini"
)

func RegisterCustomerEndpoints(m *martini.ClassicMartini) {
	m.Group("/customers", func(r martini.Router) {
		r.Get("", getCustomersHandler)
		r.Post("", createCustomerHandler)
		r.Get("/:id", getCustomerHandler)
	})

}

func getCustomersHandler(r *http.Request, w http.ResponseWriter) int {
	customers := dbcontext.GetCustomers()
	util.JSONEncode(customers, w)
	return http.StatusOK
}

func getCustomerHandler(r *http.Request, params martini.Params, w http.ResponseWriter) int {
	customer := dbcontext.GetCustomer(params["id"])
	util.JSONEncode(customer, w)
	return http.StatusOK
}

func createCustomerHandler(r *http.Request) int {
	customer := models.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	util.CheckErr(err)
	dbcontext.CreateCustomer(customer)
	return http.StatusOK
}
