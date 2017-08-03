package main

import (
	"encoding/json"
	"net/http"

	"github.com/CruizinSolutions/server/db"
	"github.com/CruizinSolutions/server/utilities"
	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()
	static := martini.Static("/TiosShopHTML")
	m.Use(static)
	m.Get("/names", getNamesHandler)
	m.Post("/create", createCustomer)
	m.Run()
}

func getNamesHandler(r *http.Request, w http.ResponseWriter) int {

	names := db.GetAllCustumers()
	js, _ := json.MarshalIndent(names, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	return http.StatusOK
}

func createCustomer(r *http.Request) int {
	customer := db.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	utilities.CheckErr(err)
	db.CreateCustomer(customer)

	return http.StatusOK
}
