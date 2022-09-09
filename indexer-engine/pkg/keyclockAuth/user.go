package keyclockauth

import (
	"fmt"

	"github.com/Nerzal/gocloak/v11"
)

type User struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
}

func (k *KeyClockInstance) CreateNewUser(realm string, user map[string]interface{}) (bool, error) {
	userData := User{
		FirstName: fmt.Sprintf("%s", user["FirstName"]),
		LastName:  fmt.Sprintf("%s", user["LastName"]),
		Username:  fmt.Sprintf("%s", user["Username"]),
		Email:     fmt.Sprintf("%s", user["Email"]),
	}

	userInKeyClockFormat := gocloak.User{
		FirstName: gocloak.StringP(userData.FirstName),
		LastName:  gocloak.StringP(userData.LastName),
		Username:  gocloak.StringP(userData.Username),
		Email:     gocloak.StringP(userData.Email),
		Enabled:   gocloak.BoolP(true),
	}
	_, err := k.client.CreateUser(k.context, k.token.AccessToken, realm, userInKeyClockFormat)
	if err != nil {
		return false, err
	}
	return true, nil
}
