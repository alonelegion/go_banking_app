package migrations

import (
	"github.com/alonelegion/go_banking_app/helpers"
	"github.com/alonelegion/go_banking_app/interfaces"
)

//type User struct {
//	gorm.Model
//	Username string
//	Email    string
//	Password string
//}
//
//type Account struct {
//	gorm.Model
//	Type    string
//	Name    string
//	Balance uint
//	UserID  uint
//}

func createAccount() {
	db := helpers.ConnectDB()
	users := &[2]interfaces.User{
		{Username: "Bruce", Email: "bruce@mail.ru"},
		{Username: "Clarke", Email: "clarke@mail.ru"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(
			10000 * int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Close()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &Account)
	defer db.Close()

	createAccount()
}

func MigrateTransactions() {
	Transactions := &interfaces.Transaction{}

	db := helpers.ConnectDB()
	db.AutoMigrate(&Transactions)
	defer db.Close()
}
