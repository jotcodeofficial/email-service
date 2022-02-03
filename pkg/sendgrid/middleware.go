package sendgrid

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// ConfigureMiddlewares will make calls to configure all the different middlewares
func configureDefaultMiddlewares(e *echo.Echo) {

	// This runs before the router
	e.Pre(middleware.RemoveTrailingSlash())

	// Use this ID to track the route through the microservices for logging, etc
	e.Pre(middleware.RequestID())

	// only enable this on certain route groups
	//e.Use(middleware.CSRF())

	//e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowCredentials, echo.HeaderCookie, echo.HeaderSetCookie},
		ExposeHeaders:    []string{echo.HeaderSetCookie},
		AllowCredentials: true,
	}))

	// If you want smaller size packets
	//e.Use(middleware.Gzip())

	// If you want to redirect all traffic to https
	//e.Use((middleware.HTTPSRedirect()))
	//e.Use((middleware.Secure()))

	// If you want the secret env to be passed in via kubernetes secrets
	//e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	// if you wish to use a casbin permission model
	// enforcer, err := casbin.NewEnforcer("casbin_auth_model.conf", "casbin_auth_policy.csv")
	// if err != nil {
	// 	fmt.Println("Error")
	// }
	// e.Use(casbin_mw.Middleware(enforcer))
}
