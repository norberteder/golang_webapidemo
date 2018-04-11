package models

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitializeDatabase(connectionString string) {
	log.Println("Initializing database ...")
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}