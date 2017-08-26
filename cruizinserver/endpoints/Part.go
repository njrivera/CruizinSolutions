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
	parts, err := dbcontext.GetParts()
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(parts, w)
}

func createPartHandler(r *http.Request, w http.ResponseWriter) {
	part := models.Part{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&part)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to add part"), w)
		return
	}
	item.Description = part.Description + " (" +
		part.Condition + ")"
	item.Type = part.Condition + " PART"
	part.ItemNum, err = dbcontext.CreateItem(item)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	err = dbcontext.CreatePart(part)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(part, w)
}

func getPartHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	part, err := dbcontext.GetPart(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(part, w)
}

func deletePartHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	err := dbcontext.DeletePart(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
}

func updatePartHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	part := models.Part{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&part)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to update part"), w)
		return
	}
	item.ItemNum = itemnum
	item.Description = part.Description + " (" +
		part.Condition + ")"
	item.Type = part.Condition + " PART"
	err = dbcontext.UpdateItem(item)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	part.ItemNum = itemnum
	err = dbcontext.UpdatePart(part)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(part, w)
}
