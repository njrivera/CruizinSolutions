package main

import (
	"os"

	"github.com/CruizinSolutions/server/dbcontext"
	"github.com/CruizinSolutions/server/endpoints"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/auth"
	_ "github.com/mattn/go-sqlite3"
)

const martiniEnv = "MARTINI_ENV"

func main() {

	m := martini.Classic()
	if os.Getenv(martiniEnv) == "production" {
		m.Use(auth.BasicFunc(func(username, password string) bool {
			valid, err := dbcontext.Authenticate(username, password)
			if err != nil {
				return false
			}
			return valid
		}))
	}

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
		endpoints.RegisterTaxEndpoint(m)
	})
	m.Run()
}
