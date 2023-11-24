package helpers

import (
	"errors"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

var secretKey = "bebekpanggang"

func GenerateToken(id uint, email string) string {
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedTOken, _ := parseToken.SignedString([]byte(secretKey))

	return signedTOken
}

func VerifyToken(c echo.Context) (jwt.MapClaims, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := c.Request().Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(secretKey), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		log.Fatal(errResponse)
		return nil, errResponse
	}

	return token.Claims.(jwt.MapClaims), nil
}
