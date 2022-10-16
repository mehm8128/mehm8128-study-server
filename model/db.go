package model

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

func InitDB() (*sqlx.DB, error) {
	user := os.Getenv("MARIADB_USERNAME")
	if user == "" {
		user = "mehm8128"
	}

	pass := os.Getenv("MARIADB_PASSWORD")
	if pass == "" {
		pass = "math8128"
	}

	host := os.Getenv("MARIADB_HOSTNAME")
	if host == "" {
		host = "localhost"
	}

	dbname := os.Getenv("MARIADB_DATABASE")
	if dbname == "" {
		dbname = "study_app"
	}
	_db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Asia", user, pass, host, dbname))
	if err != nil {
		return nil, err
	}
	db = _db

	return db, err
}
