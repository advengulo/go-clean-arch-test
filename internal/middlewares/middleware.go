package middlewares

import (
	"github.com/advengulo/go-clean-arch-test/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Register(e *echo.Echo) {
	logConfig := middleware.LoggerConfig{
		Format:           "${time_custom} | ${status} | ${method} | ${host} | ${uri} | ${user_agent} | ${latency_human}\n",
		CustomTimeFormat: "2006-01-02 15:04:05.000",
		Output:           e.Logger.Output(),
	}

	middlewareChain := []echo.MiddlewareFunc{
		middleware.LoggerWithConfig(logConfig),
		JWTAuthorization,
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
		}),
	}

	e.HTTPErrorHandler = utils.ErrorHandler
	e.Use(middlewareChain...)
}

func JWTAuthorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := utils.GetHeaderToken(c)
		if c.Request().URL.Path == "/auth/login" {
			return next(c)
		}
		if token == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization token missing")
		}
		if !utils.IsValidToken(token) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization token")
		}
		return next(c)
	}
}
