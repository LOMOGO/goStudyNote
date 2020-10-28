package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type Users struct {
	Phone string `form:"phone" json:"phone" validate:"required"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code string `form:"code" json:"code" validate:"required,len=6"`
}

func main() {
	var users = Users{
		Phone:  "18219110413",
		Passwd: "123",
		Code:   "123456",
	}
	validate := validator.New()
	err := validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
	}
}
