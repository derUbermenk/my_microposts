package data

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error

	Db, err = sql.Open("postgres", "dbname=my_microposts password=baba_yetu sslmode=disable")
	if err != nil {
		panic(err)
	}
	return
}
