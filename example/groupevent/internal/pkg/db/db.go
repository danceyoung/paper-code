package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func NewDB() *sql.DB {
	return db
}

func init() {
	tempdb, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}

	tempdb.SetMaxIdleConns(200)
	tempdb.SetMaxOpenConns(20)
	log.Println("init mysql successfully")
	db = tempdb
}
