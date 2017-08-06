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

func RegisterItemEndpoints(m *martini.ClassicMartini) {
	m.Group("/items", func(r martini.Router) {
		r.Get("", getItemsHandler)
		r.Post("", createItemHandler)
		r.Get("/:itemnum", getItemHandler)
		r.Delete("/:itemnum", deleteItemHandler)
	})
}

func getItemsHandler(r *http.Request, w http.ResponseWriter) {
	items := dbcontext.GetItems()
	util.JSONEncode(items, w)
}

func createItemHandler(r *http.Request, w http.ResponseWriter) {
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&item)
	util.CheckErr(err)
	dbcontext.CreateItem(item)
	util.JSONEncode(item, w)
}

func getItemHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	item := dbcontext.GetItem(itemnum)
	util.JSONEncode(item, w)
}

func deleteItemHandler(r *http.Request, params martini.Params) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	dbcontext.DeleteItem(itemnum)
}
