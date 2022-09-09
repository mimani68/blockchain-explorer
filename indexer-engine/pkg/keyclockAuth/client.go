package keyclockauth

import (
	"context"
	"fmt"

	"github.com/Nerzal/gocloak/v11"
)

type IAuthentication interface {
	CreateClient(ctx context.Context, username, password, realmName, authServerAddress string) error
	CreateNewUser(realm string, user map[string]interface{}) (bool, error)
}

func KeyClockBuilder() IAuthentication {
	return &KeyClockInstance{}
}

type KeyClockInstance struct {
	keyClockAddress string
	token           *gocloak.JWT
	client          gocloak.GoCloak
	context         context.Context
}

func (k *KeyClockInstance) CreateClient(ctx context.Context, username, password, realmName, authServerAddress string) error {
	k.context = ctx
	k.keyClockAddress = authServerAddress
	k.client = gocloak.NewClient(authServerAddress)
	token, err := k.client.LoginAdmin(k.context, username, password, realmName)
	if err != nil {
		fmt.Printf("[ERROR][AUTH] %s\n", err.Error())
		return err
	}
	k.token = token
	return nil
}
