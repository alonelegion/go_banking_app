package main

import (
	"github.com/alonelegion/go_banking_app/api"
	"github.com/alonelegion/go_banking_app/database"
)

func main() {
	database.InitDatabase()
	api.StartApi()
}
