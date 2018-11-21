package main

import (
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/thethan/team_picker/core"
	"github.com/thethan/team_picker/team_picker"
)


func main() {
	dependencies := mustSetupDependencies()

	router := echo.New()
	//log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
	router.Static("/", "")

	setupUnsecuredRoutes(dependencies, router)
	setupSecuredRoutes(dependencies, router)

	dependencies.Logger.Fatal(router.Start(":8080"))
}

func mustSetupDependencies() *core.Dependencies {
	// Load Configs
	config, err := core.SetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Setup dependencies
	dependencies := core.Dependencies{}
	dependencies.Config = config
	dependencies.Logger = log.New()
	dependencies.Logger.Formatter = &log.JSONFormatter{}

	//Connect to Mysql
	return &dependencies
}

func setupUnsecuredRoutes(dependencies *core.Dependencies, router *echo.Echo) {
	//	Health
	//healthGroup := router.Group("")
	//healthAPI := &health.API{
	//	Dependencies: dependencies,
	//}
	//healthAPI.SetupRoutes(healthGroup)
	catalogAPI := &team_picker.API{
		Dependencies:   dependencies,
	}
	catalogAPI.SetupTokenRoutes(router)

}

func setupSecuredRoutes(dependencies *core.Dependencies, router *echo.Echo) {
	//	Catalog
	catalogGroup := router.Group("games")

	catalogAPI := &team_picker.API{
		Dependencies:   dependencies,
	}
	catalogAPI.SetupRoutes(catalogGroup)
}