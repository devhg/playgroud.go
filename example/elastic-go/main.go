package main

import (
	"log"
	"net/http"
	"time"

	"github.com/devhg/es/bootstrap"
	"github.com/devhg/es/config"
)

// https://segmentfault.com/a/1190000024438897

func main() {

	conf := config.LoadAndInit("./config/config.yml")

	r := bootstrap.MustInit(conf)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    1000 * time.Millisecond,
		WriteTimeout:   1000 * time.Millisecond,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
