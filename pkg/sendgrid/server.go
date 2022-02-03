package sendgrid

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

// Start - Starts the program
func Start() {

	// configure viper
	startViperConfiguration()

	// Configure Middlewares
	configureDefaultMiddlewares(e)

	// Configure Routes
	configureRoutes()

	fmt.Println(config.AppEnv)
	e.Logger.Fatal(e.Start(fmt.Sprintf(config.AppEnv+":%s", config.Port)))
}

// ConfigureRoutes will make calls to configure all the different routes for fiber
func configureRoutes() {

	configureDefaultRoutes()
	configureCustomRoutes()

}
