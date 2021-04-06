package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func ProvideDBConnection() *sql.DB {
	db , err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/bank_mock")

	if err != nil {
		log.Fatal(err)
	}
	return db
}