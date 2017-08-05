package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/CruizinSolutions/cruizinserver/dbcontext"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/go-martini/martini"
)

func RegisterVehicleEndpoints(m *martini.ClassicMartini) {
	m.Group("/vehicles", func(r martini.Router) {
		r.Get("", getVehiclesHandler)
		r.Post("", createVehicleHandler)
		r.Get("/:id", getVehicleHandler)
	})

}

func getVehiclesHandler(r *http.Request, w http.ResponseWriter) int {
	vehicles := dbcontext.GetVehicles()
	util.JSONEncode(vehicles, w)
	return http.StatusOK
}

func getVehicleHandler(r *http.Request, params martini.Params, w http.ResponseWriter) int {
	vehicle := dbcontext.GetVehicle(params["id"])
	util.JSONEncode(vehicle, w)
	return http.StatusOK
}

func createVehicleHandler(r *http.Request, w http.ResponseWriter) int {
	vehicle := models.Vehicle{}
	err := json.NewDecoder(r.Body).Decode(&vehicle)
	util.CheckErr(err)
	dbcontext.CreateVehicle(vehicle)
	util.JSONEncode(vehicle, w)
	return http.StatusOK
}
