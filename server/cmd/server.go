package cmd

import (
	"log"
	"net/http"

	"github.com/solidiquis/djangolang/config/defaults"
	"github.com/solidiquis/djangolang/config/routes"
)

var (
	srv *http.Server
)

func startServer() {
	var err error

	srv := &http.Server{
		Addr:    defaults.DefPort,
		Handler: routes.Routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
