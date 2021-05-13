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

func RegisterPackageEndpoints(m *martini.ClassicMartini) {
	m.Group("/packages", func(r martini.Router) {
		r.Get("", getPackagesHandler)
		r.Post("", createPackageHandler)
		r.Get("/:itemnum", getPackageHandler)
		r.Delete("/:itemnum", deletePackageHandler)
		r.Put("/:itemnum", updatePackageHandler)
	})
}

func getPackagesHandler(r *http.Request, w http.ResponseWriter) {
	packages, err := dbcontext.GetPackages()
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(packages, w)
}

func createPackageHandler(r *http.Request, w http.ResponseWriter) {
	pack := models.Package{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&pack)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to add package"), w)
		return
	}
	item.Description = pack.Description
	item.Type = "PACKAGE"
	pack.ItemNum, err = dbcontext.CreateItem(item)
	if err != nil {
		util.JSONEncode(err, w)
	}
	err = dbcontext.CreatePackage(pack)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(pack, w)
}

func getPackageHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	pack, err := dbcontext.GetPackage(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(pack, w)
}

func deletePackageHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	err := dbcontext.DeletePackage(itemnum)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
}

func updatePackageHandler(r *http.Request, params martini.Params, w http.ResponseWriter) {
	itemnum, _ := strconv.Atoi(params["itemnum"])
	pack := models.Package{}
	item := models.Item{}
	err := json.NewDecoder(r.Body).Decode(&pack)
	if err != nil {
		log.Println(err)
		util.JSONEncode(errors.New("Unable to update package"), w)
		return
	}
	item.ItemNum = itemnum
	item.Description = pack.Description
	item.Type = "PACKAGE"
	err = dbcontext.UpdateItem(item)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	pack.ItemNum = itemnum
	err = dbcontext.UpdatePackage(pack)
	if err != nil {
		util.JSONEncode(err, w)
		return
	}
	util.JSONEncode(pack, w)
}
