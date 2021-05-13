package dbcontext

import (
	"database/sql"
	"errors"
	"log"
)

const getUser = "SELECT * FROM users WHERE username = ? AND password = ?"

func Authenticate(username string, password string) (bool, error) {
	db, err := sql.Open("sqlite3", GetDBPath())
	if err != nil {
		log.Println(err)
		return false, errors.New("Unable to get user")
	}
	defer db.Close()
	row, err := db.Query(getUser, username, password)
	if err != nil {
		log.Println(err)
		return false, errors.New("Unable to get user")
	}
	defer row.Close()

	if row.Next() {
		return true, nil
	}
	return false, nil
}
