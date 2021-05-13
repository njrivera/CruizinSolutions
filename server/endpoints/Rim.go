package endpoints

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/CruizinSolutions/server/dbcontext"
	"github.com/CruizinSolutions/server/models"
	"github.com/CruizinSolutions/server/util"
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
	rims, err := dbcontext.GetRims()
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(rims, w)
}

func createRimHandler(r *http.Request, w http.ResponseWriter) {
	rim := models.Rim{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&rim)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to add rim"), w)
		return
	}
	item.Description = rim.Brand + " " +
		rim.Model + " " +
		rim.Size + " " +
		rim.BoltPattern + " (" +
		rim.Finish + " " +
		rim.Condition + ")"
	item.Type = rim.Condition + " RIM"
	rim.ItemNum, err = dbcontext.CreateItem(item)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	err = dbcontext.CreateRim(rim)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(rim, w)
}

func getRimHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	rim, err := dbcontext.GetRim(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(rim, w)
}

func deleteRimHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	err := dbcontext.DeleteRim(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
}

func updateRimHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	rim := models.Rim{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&rim)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to update rim"), w)
		return
	}
	item.ItemNum = itemnum
	item.Description = rim.Brand + " " +
		rim.Model + " " +
		rim.Size + " " +
		rim.BoltPattern + " (" +
		rim.Finish + " " +
		rim.Condition + ")"
	item.Type = rim.Condition + " RIM"
	err = dbcontext.UpdateItem(item)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	rim.ItemNum = itemnum
	err = dbcontext.UpdateRim(rim)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(rim, w)
}

func updateRimQtyHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	qty, _ := strconv.Atoi(params["qty"])
	rim, err := dbcontext.UpdateRimQty(itemnum, qty)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(rim, w)
}
