package services

import (
    "database/sql"
    _"github.com/lib/pq"
    "log"
)

var db *sql.DB

func ConnectDB() {
    var err error
    connStr := "host=localhost port=5432 user=postgres password=yourpassword dbname=eventosDB sslmode=disable"
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to PostgreSQL successfully")
}

func GetDB() *sql.DB {
    return db
}
