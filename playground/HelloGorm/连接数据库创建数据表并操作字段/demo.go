package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name string `gorm:"not null;unique;size:64"`
	Email string `gorm:"not null;unique;size:64"`
	Password string `gorm:"not null;unique;size:64"`
}

type Tabler interface {
	TableName() string
}

// TableName 会将 User 的表名重写为 `profiles`
func (User) TableName() string {
	return "profiles"
}

func main() {
	dsn := "root:Syq@tcp(127.0.0.1:3306)/null?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	user := User{
		Name:     "lomogo",
		Email:    "dhdpat@163.com",
		Password: "Syq",
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(time.Minute * 2)
	defer sqlDB.Close()

	//自动迁移数据表
	db.AutoMigrate(&User{})
	db.Create(&user)
	fmt.Println(&user)
}
