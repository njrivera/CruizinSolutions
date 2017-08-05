package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/CruizinSolutions/cruizinserver/dbcontext"
	"github.com/CruizinSolutions/cruizinserver/models"
	"github.com/CruizinSolutions/cruizinserver/util"
	"github.com/go-martini/martini"
)

func RegisterTireEndpoints(m *martini.ClassicMartini) {
	m.Group("/tires", func(r martini.Router) {
		r.Get("", getTiresHandler)
		r.Post("", createTireHandler)
		r.Get("/:partnum", getTireHandler)
	})

}

func getTiresHandler(r *http.Request, w http.ResponseWriter) int {
	tires := dbcontext.GetTires()
	util.JSONEncode(tires, w)
	return http.StatusOK
}

func getTireHandler(r *http.Request, params martini.Params, w http.ResponseWriter) int {
	tire := dbcontext.GetTire(params["partnum"])
	util.JSONEncode(tire, w)
	return http.StatusOK
}

func createTireHandler(r *http.Request) int {
	tire := models.Tire{}
	err := json.NewDecoder(r.Body).Decode(&tire)
	util.CheckErr(err)
	dbcontext.CreateTire(tire)
	return http.StatusOK
}
