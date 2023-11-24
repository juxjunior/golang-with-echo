package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"dbgo/config"
	"dbgo/model"
)

// CreateEmployee return with created data in database and show the data in json
// CreateEmployee godoc
// @Summary Create a new Employee
// @description Create a new employee with the input paylod
// @Tags items
// @Accept json
// @Produce json
// @Param models,Item body model.Employee true "Create employee"
// @Success 200 {object} model.Employee
// @Router /emp [post]
func CreateEmployee(c echo.Context) error {
	db := config.GetDB()

	emp := model.Employee{}

	if err := c.Bind(&emp); err != nil {
		return err
	}

	db.Debug().Create(&emp)

	fmt.Println("CreateEmployee")

	return c.JSON(http.StatusOK, emp)
}

func UpdateEmployee(c echo.Context) error {
	db := config.GetDB()

	emp := model.Employee{}

	if err := c.Bind(&emp); err != nil {
		return err
	}

	db.Model(&emp).Where("id = ?", emp.ID).Updates(model.Employee{
		Full_name: emp.Full_name,
		Age:       emp.Age,
		Email:     emp.Email,
		Division:  emp.Division,
	})

	fmt.Println("UpdateEmployee")

	return c.JSON(http.StatusOK, emp)
}

func DeleteEmployee(c echo.Context) error {
	db := config.GetDB()

	emp := model.Employee{}

	delResp := model.DeleteStruct{
		Status:  http.StatusOK,
		Message: "Berhasil Ngapus Gaes",
	}

	if err := c.Bind(&emp); err != nil {
		return err
	}

	paramId := c.Param("id")

	db.Model(&emp).Where("id = ?", paramId).Delete(&emp)

	fmt.Println("DeleteEmployee")

	return c.JSON(http.StatusOK, delResp)
}

func CreateItem(c echo.Context) error {
	db := config.GetDB()
	item := model.Item{}

	userData, ok := c.Get("userData").(jwt.MapClaims)
	if !ok {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"data":    userData,
			"message": "Failed to get data",
		})
	}

	userId := uint(userData["id"].(float64))

	if err := c.Bind(&item); err != nil {
		return err
	}

	item.EmployeeId = int(userId)

	db.Debug().Create(&item)

	fmt.Println("CreateItem")
	return c.JSON(http.StatusOK, item)
}

func Index(c echo.Context) error {
	tmpl :=
		template.Must(template.ParseGlob("template/*.html"))
	type M map[string]interface{}

	data := make(M)
	data["csrf_token"] = c.Get(config.CSRFKey)
	return tmpl.Execute(c.Response(), data)

}

func SayHello(c echo.Context) error {
	data := make(map[string]interface{})

	if err := c.Bind(&data); err != nil {
		return err
	}

	message := fmt.Sprintf("Hello %s, My Gender %s", data["name"], data["gender"])

	return c.JSON(http.StatusOK, message)
}
