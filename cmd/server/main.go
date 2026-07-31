// @title Nova ERP API
// @version 1.0
// @description Nova ERP REST API
// @BasePath /api/v1
// @host localhost:8080

package main

import (
	"log"

	_ "github.com/swaggo/files"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tamim1715/novaerp/docs"
	"github.com/tamim1715/novaerp/internal/bootstrap"
)

func main() {

	router, application := bootstrap.Bootstrap()

	defer application.Logger.Sync()

	log.Printf("%s started on port %s",
		application.Config.AppName,
		application.Config.Port,
	)

	// 2. REGISTER THE SWAGGER ROUTE
	router.GET(
		"/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler),
	)

	if err := router.Run(":" + application.Config.Port); err != nil {
		log.Fatal(err)
	}
}
