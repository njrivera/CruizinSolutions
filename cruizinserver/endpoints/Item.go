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

func RegisterItemEndpoints(m *martini.ClassicMartini) {
	m.Group("/items", func(r martini.Router) {
		r.Get("", getItemsHandler)
		r.Post("", createItemHandler)
		r.Get("/:itemnum", getItemHandler)
		r.Delete("/:itemnum", deleteItemHandler)
	})
}

func getItemsHandler(r *http.Request, w http.ResponseWriter) {
	items, err := dbcontext.GetItems()
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(items, w)
}

func createItemHandler(r *http.Request, w http.ResponseWriter) {
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to add item"), w)
		return
	}
	dbcontext.CreateItem(item)
	util.JSONEncode(item, w)
}

func getItemHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	item, err := dbcontext.GetItem(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(item, w)
}

func deleteItemHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	err := dbcontext.DeleteItem(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
}
