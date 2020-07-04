package migrations

import (
	"github.com/alonelegion/go_banking_app/database"
	"github.com/alonelegion/go_banking_app/helpers"
	"github.com/alonelegion/go_banking_app/interfaces"
)

func createAccount() {
	users := &[2]interfaces.User{
		{Username: "Bruce", Email: "bruce@mail.ru"},
		{Username: "Clarke", Email: "clarke@mail.ru"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(
			10000 * int(i+1)), UserID: user.ID}
		database.DB.Create(&account)
	}
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	Transactions := &interfaces.Transaction{}
	database.DB.AutoMigrate(&User, &Account, &Transactions)

	createAccount()
}
