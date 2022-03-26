package main

import (
	"fmt"

	"mehm8128-study-server3/model"

	"mehm8128-study-server3/router"
)

func main() {
	db, err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	fmt.Println(db) //todo:あとでセッション関連でdbを使う
	router.SetRouting()
}
