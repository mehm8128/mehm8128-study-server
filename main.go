package main

import (
	"fmt"
	"mehm8128_study_server/model"
	"mehm8128_study_server/router"
)

func main() {
	db, err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	fmt.Println(db) //todo:あとでセッション関連でdbを使う
	router.SetRouting()
}
