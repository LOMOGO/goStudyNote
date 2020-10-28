package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name   string `form:"name" validate:"required,CustomValidationErrors"`
	Age    uint8  `form:"age" validate:"gt=18"`
	Passwd string `form:"passwd" validate:"required,max=16,min=8"`
}

func main() {
	user := User{
		Name:   "司云起",
		Age:    17,
		Passwd: "11213",
	}
	validate := validator.New()
	//注册自定义验证函数
	err := validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors)
	if err != nil {
		fmt.Println(err)
	}
	err = validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
	}
}

func CustomValidationErrors(fl validator.FieldLevel) bool {
	return fl.Field().String() == "admin"
}
