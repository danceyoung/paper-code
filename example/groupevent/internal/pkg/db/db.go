package db

import (
	"database/sql"
	"log"

	"paper-code/example/groupevent/internal/pkg/cfg"

	_ "github.com/go-sql-driver/mysql"
)

// How errors are handled:
//NOT allowed to panic an application.
//NOT allowed to wrap errors.
//Return only root cause error values.

var db *sql.DB

func NewDB() *sql.DB {
	return db
}

func init() {
	tempdb, err := sql.Open("mysql", cfg.String("mysql"))
	if err != nil {
		log.Println(err)
		return
	}

	tempdb.SetMaxIdleConns(200)
	tempdb.SetMaxOpenConns(20)
	log.Println("init mysql successfully")
	db = tempdb
}
