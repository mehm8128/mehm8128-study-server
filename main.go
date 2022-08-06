package main

import (
	"fmt"
	"mehm8128_study_server/model"
	"mehm8128_study_server/router"
)

func main() {
	_, err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("DB Error: %w", err))
	}
	router.SetRouting()
}
