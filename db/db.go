package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	//本番環境では実行されないようにする
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected")
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, _ := db.DB()
	if err := dbSQL.Close(); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Closed")
}
