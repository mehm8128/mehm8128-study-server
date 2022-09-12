package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

func InitDB() (*sqlx.DB, error) {
	_db, err := sqlx.Open("mariadb", "user=mehm8128 password=math8128 dbname=study_app sslmode=disable")
	//_db, err := sqlx.Open("mariadb", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	db = _db

	return db, err
}
