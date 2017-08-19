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

func RegisterVehicleEndpoints(m *martini.ClassicMartini) {
	m.Group("/vehicles", func(r martini.Router) {
		r.Get("", getVehiclesHandler)
		r.Post("", createVehicleHandler)
		r.Get("/:vid", getVehicleHandler)
		r.Delete("/:vid", deleteVehicleHandler)
		r.Put("/:vid", updateVehicleHandler)
	})
}

func getVehiclesHandler(r *http.Request, w http.ResponseWriter) {
	vehicles, err := dbcontext.GetVehicles()
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(vehicles, w)
}

func createVehicleHandler(r *http.Request, w http.ResponseWriter) {
	vehicle := models.Vehicle{}
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to add vehicle"), w)
		return
	}
	err = dbcontext.CreateVehicle(vehicle)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(vehicle, w)
}

func getVehicleHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	vid, _ := strconv.Atoi(params["vid"])
	vehicle, err := dbcontext.GetVehicle(vid)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(vehicle, w)
}

func deleteVehicleHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	vid, _ := strconv.Atoi(params["vid"])
	err := dbcontext.DeleteVehicle(vid)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
}

func updateVehicleHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	vid, _ := strconv.Atoi(params["vid"])
	vehicle := models.Vehicle{}
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to update vehicle"), w)
		return
	}
	vehicle.Vid = vid
	err = dbcontext.UpdateVehicle(vehicle)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(vehicle, w)
}
