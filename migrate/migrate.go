package main

import (
	"fmt"
	"go_rest_api/db"
	"go_rest_api/model"
	"os"
)

func main() {
	//環境変数を設定
	os.Setenv("GO_ENV", "dev") //本番環境ではコメントアウトしたほうがいい？

	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)

	// DBにテーブルを作成
	dbConn.AutoMigrate(&model.User{}, &model.Work{})
}
