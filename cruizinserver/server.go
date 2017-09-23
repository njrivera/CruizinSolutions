package main

import (
	"github.com/CruizinSolutions/cruizinserver/endpoints"
	"github.com/go-martini/martini"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	m := martini.Classic()
	//
	// Uncomment for production build
	//
	/*m.Use(auth.BasicFunc(func(username, password string) bool {
		valid, err := dbcontext.Authenticate(username, password)
		if err != nil {
			return false
		}
		return valid
	}))*/
	m.Group("/api", func(r martini.Router) {
		endpoints.RegisterCustomerEndpoints(m)
		endpoints.RegisterItemEndpoints(m)
		endpoints.RegisterOrderEndpoints(m)
		endpoints.RegisterPartEndpoints(m)
		endpoints.RegisterRimEndpoints(m)
		endpoints.RegisterServiceEndpoints(m)
		endpoints.RegisterPackageEndpoints(m)
		endpoints.RegisterTireEndpoints(m)
		endpoints.RegisterVehicleEndpoints(m)
		endpoints.RegisterReportEndpoints(m)
	})
	m.Run()
}
