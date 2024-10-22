package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"pismo/app"
	"pismo/cmd/api"
)

// @title         Pismo Test API
// @version       1.0
// @description   Implementation of Pismo Test API.
// @contact.name  Tcharles Pereira da Silva
// @contact.url   https://tcharlez.in
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {

	app.Application.Boot()
	defer app.Application.Shutdown()

	webPort := os.Getenv("WEB_PORT")

	// Define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: api.Routes(),
	}

	// Start server
	app.Application.Log.Info(fmt.Sprintf("Starting webserver at port %s", webPort))

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic("[Server] ", err)
	}
}
