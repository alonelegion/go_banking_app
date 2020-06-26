package users

import (
	"github.com/alonelegion/go_banking_app/helpers"
	"github.com/alonelegion/go_banking_app/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func prepareToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount, withToken bool) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Account:  accounts,
	}

	var response = map[string]interface{}{
		"message": "all is fine",
	}
	if withToken {
		var token = prepareToken(user)
		response["jwt"] = token
	}
	response["data"] = responseUser

	return response
}

func Login(username string, pass string) map[string]interface{} {
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		db := helpers.ConnectDB()
		user := &interfaces.User{}
		if db.Where("username = ?", username).First(&user).RecordNotFound() {
			return map[string]interface{}{
				"message": "User not found",
			}
		}

		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{
				"message": "Wrong password",
			}
		}

		accounts := []interfaces.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		defer db.Close()

		var response = prepareResponse(user, accounts, true)
		return response
	} else {
		return map[string]interface{}{
			"message": "not valid values",
		}
	}

}

func Register(username string, email string, pass string) map[string]interface{} {
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		db := helpers.ConnectDB()
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{Username: username, Email: email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(username + "'s" + " account"), Balance: 0, UserID: user.ID}
		db.Create(&account)

		defer db.Close()

		accounts := []interfaces.ResponseAccount{}
		respAccount := interfaces.ResponseAccount{
			ID:      account.ID,
			Name:    account.Name,
			Balance: int(account.Balance),
		}
		accounts = append(accounts, respAccount)
		var response = prepareResponse(user, accounts, true)

		return response
	} else {
		return map[string]interface{}{
			"message": "not valid values",
		}
	}
}

func GetUser(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)

	if isValid {
		db := helpers.ConnectDB()

		user := &interfaces.User{}
		if db.Where("id = ?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{
				"message": "User not found",
			}
		}

		accounts := []interfaces.ResponseAccount{}
		db.Table("accounts").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

		defer db.Close()

		var response = prepareResponse(user, accounts, false)
		return response
	} else {
		return map[string]interface{}{
			"message": "Not valid token",
		}
	}
}
