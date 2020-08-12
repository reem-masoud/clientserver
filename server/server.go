package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/clientserver/configs"
	"github.com/clientserver/routes"
)

// StartServer starts the http server
func StartServer() {
	e := routes.Router()

	appCfg := configs.App()

	portStr := fmt.Sprintf(":%d", appCfg.Port)

	server := &http.Server{
		Addr: portStr,

		Handler: e,
	}

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Printf("Listening on port::%d...\n", appCfg.Port)
	<-sigs
}
