package app

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/go-crud-category-api/helper"
)

func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:#21012123Op@tcp(localhost:3306)/crud_category")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
