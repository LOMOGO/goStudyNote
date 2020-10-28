package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
)

type User struct {
	Name string `form:"name" validate:"required,CustomValidationErrors" label:"姓名"`
	Age uint8 `form:"age" validate:"required,gt=18" label:"年龄"`
	Passwd string `form:"passwd" validate:"required,max=20,min=6"`
}

func main() {
	user := User{
		Name:   "admin",
		Age:    21,
		Passwd: "123345678",
	}
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	//验证器注册翻译器
	err := zh_trans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatal(err)
	}
	//注册自定义函数
	_=validate.RegisterValidation("CustomValidationErrors", CustomValidationErrors)
	//注册一个标签命名函数，用于获取struct tag里面自定义的label作为字段名
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("label")
		return name
	})
	//根据提供的标记注册翻译
	validate.RegisterTranslation("CustomValidationErrors", trans, func(ut ut.Translator) error {
		return ut.Add("CustomValidationErrors", "{0}错误", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("CustomValidationErrors", fe.Field(), fe.Field())
		return t
	})
	err = validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err)
		}
	}
}

func CustomValidationErrors(fl validator.FieldLevel) bool {
	return fl.Field().String() == "司云起"
}
