package database

import (
	"github.com/alonelegion/go_banking_app/helpers"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	database, err := gorm.Open("postgres", "host=database-1.c1fzlyaabojw.eu-central-1.rds.amazonaws.com port=5432 user=postgres dbname=bankapp password=zw345b7u sslmode=disable")
	helpers.HandleErr(err)

	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)
	DB = database
}
