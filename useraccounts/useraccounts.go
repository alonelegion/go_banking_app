package useraccounts

import (
	"github.com/alonelegion/go_banking_app/helpers"
	"github.com/alonelegion/go_banking_app/interfaces"
)

func updateAccount(id uint, amount int) {
	db := helpers.ConnectDB()
	db.Model(&interfaces.Account{}).Where("id = ?", id).Update("balance", amount)
	defer db.Close()
}
