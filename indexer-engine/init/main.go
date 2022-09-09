package main

import "app.io/cmd"

// @title Crypto Indexer
// @version 1.0
// @description Server side Crypto Indexer
// @termsOfService http://app.io/api
//
// @contact.name API Support
// @contact.url http://app.io/support
// @contact.email support@app.io
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host app.io:3000
// @BasePath /api
func main() {
	cmd.Execute()
}
