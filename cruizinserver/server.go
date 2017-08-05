package main

import (
	"github.com/CruizinSolutions/cruizinserver/endpoints"
	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()

	endpoints.RegisterCustomerEndpoints(m)
	endpoints.RegisterServiceEndpoints(m)
	endpoints.RegisterTireEndpoints(m)
	endpoints.RegisterVehicleEndpoints(m)

	m.Run()
}
