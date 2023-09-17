package main

import (
	"fmt"
	"todo-app/db"
	"todo-app/model"
)

func main() {
	dbConn := db.NewDb()
	defer fmt.Println("Successfully Migration")
	defer db.CloseDb(dbConn)
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		return
	}
}
