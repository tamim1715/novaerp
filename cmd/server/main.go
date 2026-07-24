package main

import (
	"log"

	"github.com/tamim1715/novaerp/internal/bootstrap"
)

func main() {

	router, application := bootstrap.Bootstrap()

	defer application.Logger.Sync()

	log.Printf("%s started on port %s",
		application.Config.AppName,
		application.Config.Port,
	)

	if err := router.Run(":" + application.Config.Port); err != nil {
		log.Fatal(err)
	}
}
