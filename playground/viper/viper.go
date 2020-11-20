package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

type Database struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Dbname   string `mapstructure:"dbname"`
}

var DatabaseConf Database

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("配置信息改变:", e.Name)
	})
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("数据读取失败%v", err)
	}
	err = viper.UnmarshalKey("mysql", &DatabaseConf)
	if err != nil {
		log.Printf("配置数据绑定失败：%v", err)
	}
	fmt.Println(DatabaseConf.Dbname)
	fmt.Println("每天一遍，快乐再见")


}
