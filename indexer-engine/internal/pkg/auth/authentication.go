package auth

import (
	"context"
	"fmt"

	auth "app.io/pkg/keyclockAuth"
)

type authentication struct {
	PluginName     string
	keyClockPlugin auth.KeyClockInstance
	casbinPlugin   auth.KeyClockInstance
	cognitoPlugin  auth.KeyClockInstance
	oauthPlugin    auth.KeyClockInstance
}

type IAuthentication interface {
	CreateClient(ctx context.Context, username, password, realmName, authServerAddress string) error
	CreateNewUser(realm string, user map[string]interface{}) (bool, error)
}

type authClass struct {
	pluginName string
	plugin     IAuthentication
}

func AuthPluginControler() *authClass {
	return &authClass{}
}

func (a *authClass) LoadPlugin(pulignName string, pluginInstance IAuthentication) {
	a.pluginName = pulignName
	a.plugin = pluginInstance
}

func (a *authClass) GetInstance() IAuthentication {
	return a.plugin
}

func (a *authClass) SetParameter(paramName, valur string) {
	fmt.Println("Not implemented yet")
}
