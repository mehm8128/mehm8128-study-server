package main

import (
	"fmt"

	"github.com/mehm8128/mehm8128-study-server/model"
	"github.com/mehm8128/mehm8128-study-server/router"
)

func main() {
	db, err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	fmt.Println(db) //todo:あとでセッション関連でdbを使う
	router.SetRouting()
}
