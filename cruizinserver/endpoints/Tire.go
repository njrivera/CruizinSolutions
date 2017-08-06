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

func RegisterTireEndpoints(m *martini.ClassicMartini) {
	m.Group("/tires", func(r martini.Router) {
		r.Get("", getTiresHandler)
		r.Post("", createTireHandler)
		r.Get("/:itemnum", getTireHandler)
		r.Delete("/:itemnum", deleteTireHandler)
	})
}

func getTiresHandler(r *http.Request, w http.ResponseWriter) {
	tires := dbcontext.GetTires()
	util.JSONEncode(tires, w)
}

func createTireHandler(r *http.Request, w http.ResponseWriter) {
	tire := models.Tire{}
	err := json.NewDecoder(r.Body).Decode(&tire)
	util.CheckErr(err)
	dbcontext.CreateTire(tire)
	util.JSONEncode(tire, w)
}

func getTireHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	tire := dbcontext.GetTire(itemnum)
	util.JSONEncode(tire, w)
}

func deleteTireHandler(r *http.Request, params martini.Params) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	dbcontext.DeleteTire(itemnum)
}
