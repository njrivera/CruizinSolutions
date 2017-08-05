package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/CruizinSolutions/cruizinserver/dbcontext"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/go-martini/martini"
)

func RegisterServiceEndpoints(m *martini.ClassicMartini) {
	m.Group("/services", func(r martini.Router) {
		r.Get("", getServicesHandler)
		r.Post("", createServiceHandler)
		r.Get("/:partnum", getServiceHandler)
	})

}

func getServicesHandler(r *http.Request, w http.ResponseWriter) int {
	services := dbcontext.GetServices()
	util.JSONEncode(services, w)
	return http.StatusOK
}

func getServiceHandler(r *http.Request, params martini.Params, w http.ResponseWriter) int {
	service := dbcontext.GetService(params["partnum"])
	util.JSONEncode(service, w)
	return http.StatusOK
}

func createServiceHandler(r *http.Request) int {
	service := models.Service{}
	err := json.NewDecoder(r.Body).Decode(&service)
	util.CheckErr(err)
	dbcontext.CreateService(service)
	return http.StatusOK
}
