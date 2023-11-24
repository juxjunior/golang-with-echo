package middlerware

import (
	"dbgo/helpers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authentication() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			verifyToken, err := helpers.VerifyToken(c)
			if err != nil {
				log.Fatalf(err.Error())
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Unauthorized",
					"message": "token error",
				})
			}

			c.Set("userData", verifyToken)

			return next(c)
		}
	}
}
