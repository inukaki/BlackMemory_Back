package main

import (
	"fmt"
	"go_rest_api/db"
	"go_rest_api/model"
	"os"
)

func main() {
	os.Setenv("GO_ENV", "dev")
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Task{})
}
