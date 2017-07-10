package main

import (
	"log"
	"net/http"

	environment "github.com/hellomd/go-microservice-starter/environment"
	"github.com/hellomd/go-microservice-starter/router"
	"github.com/hellomd/go-sdk/config"
)

func main() {
	handler, err := environment.NewDevelopmentEnv().GetHandler()
	if err != nil {
		log.Fatal("Failed to initialize handler ", err)
	}

	handler.UseHandler(router.NewRouter())

	log.Fatal(http.ListenAndServe(":"+config.Get(environment.PortCfgKey), handler))
}
