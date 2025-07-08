package app

import (
	"customer-restful-api/helper"
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewDB() *sql.DB {
	connString := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connString)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
