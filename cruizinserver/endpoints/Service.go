package endpoints

import (
	"encoding/json"
	"errors"
	"log"
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
	services, err := dbcontext.GetServices()
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(services, w)
}

func createServiceHandler(r *http.Request, w http.ResponseWriter) {
	service := models.Service{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to add service"), w)
		return
	}
	item.Description = service.Description
	service.ItemNum, err = dbcontext.CreateItem(item)
	if err != nil {
		util.JSONEncode(err, w)
	}
	err = dbcontext.CreateService(service)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(service, w)
}

func getServiceHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	service, err := dbcontext.GetService(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(service, w)
}

func deleteServiceHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	err := dbcontext.DeleteService(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
}

func updateServiceHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	service := models.Service{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to update service"), w)
		return
	}
	item.Description = service.Description
	err = dbcontext.UpdateItem(itemnum, item.Description)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	service.ItemNum = itemnum
	err = dbcontext.UpdateService(service)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(service, w)
}
