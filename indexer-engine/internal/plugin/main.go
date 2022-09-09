package plugin

import (
	"context"
	"fmt"

	"app.io/internal/pkg/auth"
	authPkg "app.io/pkg/keyclockAuth"
)

func LoadAllPlugins() (error, auth.IAuthentication) {
	// Prepare auth mechanisem
	authPlugin := auth.AuthPluginControler()
	// you can inject any auth tools/pkg
	keyclockPkg := authPkg.KeyClockBuilder()
	authPlugin.LoadPlugin("keyclock", keyclockPkg)
	authPluginInstance := authPlugin.GetInstance()
	err := authPluginInstance.CreateClient(context.Background(), "mahdi", "1234", "appRealm", "https://keyclock.local")
	if err != nil {
		fmt.Println("[AUTH] error while trying initializing AUTH PLUGIN")
	}
	return nil, authPluginInstance
}
