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

func RegisterPartEndpoints(m *martini.ClassicMartini) {
	m.Group("/parts", func(r martini.Router) {
		r.Get("", getPartsHandler)
		r.Post("", createPartHandler)
		r.Get("/:itemnum", getPartHandler)
		r.Delete("/:itemnum", deletePartHandler)
		r.Put("/:itemnum", updatePartHandler)
	})
}

func getPartsHandler(r *http.Request, w http.ResponseWriter) {
	parts := dbcontext.GetParts()
	util.JSONEncode(parts, w)
}

func createPartHandler(r *http.Request, w http.ResponseWriter) {
	part := models.Part{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&part)
	util.CheckErr(err)
	item.Description = part.Description + " (" +
		part.Condition + ")"
	part.ItemNum = dbcontext.CreateItem(item)
	dbcontext.CreatePart(part)
	util.JSONEncode(part, w)
}

func getPartHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	part := dbcontext.GetPart(itemnum)
	util.JSONEncode(part, w)
}

func deletePartHandler(r *http.Request, params martini.Params) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	dbcontext.DeletePart(itemnum)
}

func updatePartHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	part := models.Part{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&part)
	util.CheckErr(err)
	item.Description = part.Description + " (" +
		part.Condition + ")"
	part.ItemNum = itemnum
	dbcontext.UpdateItem(itemnum, item.Description)
	dbcontext.UpdatePart(part)
	util.JSONEncode(part, w)
}
