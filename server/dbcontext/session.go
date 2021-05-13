package dbcontext

import (
	"os"
	"sync"
)

const (
	dbPathEnv  = "CRUIZIN_DB_PATH"
	martiniEnv = "MARTINI_ENV"
)

var (
	once   sync.Once
	DbPath string
)

func GetDBPath() string {
	DbPath := os.Getenv(dbPathEnv)
	if DbPath == "" {
		if os.Getenv(martiniEnv) == "production" {
			DbPath = "/opt/CruizinSolutions/CruizinSolutions.db"
		} else {
			DbPath = "./database/CruizinSolutions.db"
		}
	}
	return DbPath
}
