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

func RegisterVehicleEndpoints(m *martini.ClassicMartini) {
	m.Group("/vehicles", func(r martini.Router) {
		r.Get("", getVehiclesHandler)
		r.Post("", createVehicleHandler)
		r.Get("/:vid", getVehicleHandler)
		r.Delete("/:vid", deleteVehicleHandler)
	})
}

func getVehiclesHandler(r *http.Request, w http.ResponseWriter) {
	vehicles := dbcontext.GetVehicles()
	util.JSONEncode(vehicles, w)
}

func createVehicleHandler(r *http.Request, w http.ResponseWriter) {
	vehicle := models.Vehicle{}
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	util.CheckErr(err)
	dbcontext.CreateVehicle(vehicle)
	util.JSONEncode(vehicle, w)
}

func getVehicleHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	vid, _ := strconv.Atoi(params["vid"])
	vehicle := dbcontext.GetVehicle(vid)
	util.JSONEncode(vehicle, w)
}

func deleteVehicleHandler(r *http.Request, params martini.Params) {
	vid, _ := strconv.Atoi(params["vid"])
	dbcontext.DeleteVehicle(vid)
}
