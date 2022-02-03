package sendgrid

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ConfigureDefaultRoutes - Configure all the default routes here
func configureDefaultRoutes() {

	e.GET("/", defaultRoute)

}

func defaultRoute(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, "the email service is running")
}
