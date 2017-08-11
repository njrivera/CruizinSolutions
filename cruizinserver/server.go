package main

import (
	"github.com/CruizinSolutions/cruizinserver/endpoints"
	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()
	m.Group("/api", func(r martini.Router) {
		endpoints.RegisterCustomerEndpoints(m)
		endpoints.RegisterItemEndpoints(m)
		endpoints.RegisterTireEndpoints(m)
		endpoints.RegisterVehicleEndpoints(m)
	})
	m.Run()
}
