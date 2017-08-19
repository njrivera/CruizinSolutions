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

func RegisterTireEndpoints(m *martini.ClassicMartini) {
	m.Group("/tires", func(r martini.Router) {
		r.Get("", getTiresHandler)
		r.Post("", createTireHandler)
		r.Get("/:itemnum", getTireHandler)
		r.Delete("/:itemnum", deleteTireHandler)
		r.Put("/:itemnum", updateTireHandler)
		r.Put("/:itemnum/:qty", updateTireQtyHandler)
	})
}

func getTiresHandler(r *http.Request, w http.ResponseWriter) {
	tires, err := dbcontext.GetTires()
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(tires, w)
}

func createTireHandler(r *http.Request, w http.ResponseWriter) {
	tire := models.Tire{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&tire)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to add tire"), w)
		return
	}
	item.Description = tire.Brand + " " +
		tire.Model + " " +
		tire.Size + " " +
		tire.ServiceDesc + " (" +
		tire.Condition + ")"
	tire.ItemNum, err = dbcontext.CreateItem(item)
	if err != nil {
		util.JSONEncode(err, w)
	}
	err = dbcontext.CreateTire(tire)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(tire, w)
}

func getTireHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	tire, err := dbcontext.GetTire(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(tire, w)
}

func deleteTireHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	err := dbcontext.DeleteTire(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
}

func updateTireHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	tire := models.Tire{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&tire)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to update tire"), w)
		return
	}
	item.Description = tire.Brand + " " +
		tire.Model + " " +
		tire.Size + " " +
		tire.ServiceDesc + " (" +
		tire.Condition + ")"
	err = dbcontext.UpdateItem(itemnum, item.Description)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	tire.ItemNum = itemnum
	err = dbcontext.UpdateTire(tire)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(tire, w)
}

func updateTireQtyHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	qty, _ := strconv.Atoi(params["qty"])
	tire, err := dbcontext.UpdateTireQty(itemnum, qty)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(tire, w)
}
