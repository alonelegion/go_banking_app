package transactions

import (
	"github.com/alonelegion/go_banking_app/helpers"
	"github.com/alonelegion/go_banking_app/interfaces"
)

func CreateTransaction(From uint, To uint, Amount int) {
	db := helpers.ConnectDB()
	transactions := &interfaces.Transaction{From: From, To: To, Amount: Amount}
	db.Create(&transactions)

	defer db.Close()
}
