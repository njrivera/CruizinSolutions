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

func RegisterRimEndpoints(m *martini.ClassicMartini) {
	m.Group("/rims", func(r martini.Router) {
		r.Get("", getRimsHandler)
		r.Post("", createRimHandler)
		r.Get("/:itemnum", getRimHandler)
		r.Delete("/:itemnum", deleteRimHandler)
		r.Put("/:itemnum", updateRimHandler)
		r.Put("/:itemnum/:qty", updateRimQtyHandler)
	})
}

func getRimsHandler(r *http.Request, w http.ResponseWriter) {
	rims := dbcontext.GetRims()
	util.JSONEncode(rims, w)
}

func createRimHandler(r *http.Request, w http.ResponseWriter) {
	rim := models.Rim{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&rim)
	util.CheckErr(err)
	item.Description = rim.Brand + " " +
		rim.Model + " " +
		rim.Size + " " +
		rim.BoltPattern + " (" +
		rim.Finish + " " +
		rim.Condition + ")"
	rim.ItemNum = dbcontext.CreateItem(item)
	dbcontext.CreateRim(rim)
	util.JSONEncode(rim, w)
}

func getRimHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	rim := dbcontext.GetRim(itemnum)
	util.JSONEncode(rim, w)
}

func deleteRimHandler(r *http.Request, params martini.Params) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	dbcontext.DeleteRim(itemnum)
}

func updateRimHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	rim := models.Rim{}
	err := json.NewDecoder(r.Body).Decode(&rim)
	util.CheckErr(err)
	rim.ItemNum = itemnum
	dbcontext.UpdateRim(rim)
	util.JSONEncode(rim, w)
}

func updateRimQtyHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	qty, _ := strconv.Atoi(params["qty"])
	rim := dbcontext.UpdateRimQty(itemnum, qty)
	util.JSONEncode(rim, w)
}
