package main

import (
	"fmt"
	"mehm8128_study_server/model"
	"mehm8128_study_server/router"

	"github.com/srinathgs/mysqlstore"
)

func main() {
	db, err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}

	store, err := mysqlstore.NewMySQLStoreFromConnection(db.DB, "sessions", "/", 3600, []byte("secret-token"))
	if err != nil {
		panic(err)
	}
	router.SetRouting(store)
}
