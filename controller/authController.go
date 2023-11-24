package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"dbgo/config"
	"dbgo/helpers"
	"dbgo/model"
)

func Login(c echo.Context) error {
	db := config.GetDB()
	Emp := model.Employee{}
	password := ""

	if err := c.Bind(&Emp); err != nil {
		return err
	}

	password = Emp.Password

	if err := db.Debug().Where("email = ?", Emp.Email).Take(&Emp).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized",
			"message": "invalid email/pass",
		})
	}

	comparePass := helpers.ComparePass([]byte(Emp.Password), []byte(password))

	if !comparePass {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   "Unauthorized - From Compare Pass",
			"message": "invalid email/pass",
		})
	}

	token := helpers.GenerateToken(uint(Emp.ID), Emp.Email)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
