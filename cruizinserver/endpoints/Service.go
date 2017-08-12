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

func RegisterServiceEndpoints(m *martini.ClassicMartini) {
	m.Group("/services", func(r martini.Router) {
		r.Get("", getServicesHandler)
		r.Post("", createServiceHandler)
		r.Get("/:itemnum", getServiceHandler)
		r.Delete("/:itemnum", deleteServiceHandler)
		r.Put("/:itemnum", updateServiceHandler)
	})
}

func getServicesHandler(r *http.Request, w http.ResponseWriter) {
	services := dbcontext.GetServices()
	util.JSONEncode(services, w)
}

func createServiceHandler(r *http.Request, w http.ResponseWriter) {
	service := models.Service{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&service)
	util.CheckErr(err)
	item.Description = service.Description
	item.Taxable = "true"
	service.ItemNum = dbcontext.CreateItem(item)
	dbcontext.CreateService(service)
	util.JSONEncode(service, w)
}

func getServiceHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	service := dbcontext.GetService(itemnum)
	util.JSONEncode(service, w)
}

func deleteServiceHandler(r *http.Request, params martini.Params) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	dbcontext.DeleteService(itemnum)
}

func updateServiceHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	service := models.Service{}
	err := json.NewDecoder(r.Body).Decode(&service)
	util.CheckErr(err)
	service.ItemNum = itemnum
	dbcontext.UpdateService(service)
	util.JSONEncode(service, w)
}
