package model

import (
	"dbgo/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type M map[string]interface{}

type Item struct {
	Name       string `json:"name" form:"name`
	EmployeeId int
	Employee   *Employee
}

type Employee struct {
	ID        int    `json:"id" form:"id" swagger:"ID(Employee ID)" example: "1"`
	Full_name string `json:"full_name" form:"full_name" valid:"required" swagger:"Full Name(Employee ID)" example: "Kuru"`
	Email     string `json:"email" form:"email" valid:"required" swagger:"Email(Employee email)" example: "hehe@gmail.com"`
	Password  string `json:"password" form:"password" valid:"required" swagger:"Passrowd(Employee password) example: "19jdiojw0398ueq3ojf"`
	Age       int    `json:"age" form:"age" valid:"required" swagger:"Age(Employee Age)" example: "20"`
	Division  string `json:"division" form:"division" valid:"required" swagger:"Division(Employee Division)" example: "IT Dev"`
	Item      Item
}

type DeleteStruct struct {
	Status  int    `json:"status" form:"status"`
	Message string `json:"message" form:"message"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(e)

	if err != nil {
		err = errCreate
		return
	}

	e.Password = helpers.HashPass(e.Password)
	err = nil
	return
}
