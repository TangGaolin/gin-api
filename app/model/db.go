package models

import (
	"fmt"
	"gin-api/pkg/logs"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var db *gorm.DB

//连接池数量
const MAX_OPEN_POOL_SIZE = 5
const MAX_IDLE_SIZE = 3

func Init() {
	var err error
	db, err = GetConnect()
	if err != nil {
		logs.Logger.Error("gorm.Open err:" + err.Error())
		panic(err)
	}
	db.SingularTable(true) //实现结构体名为非复数形式

	//设置连接池配置
	db.DB().SetMaxOpenConns(MAX_OPEN_POOL_SIZE)
	db.DB().SetMaxIdleConns(MAX_IDLE_SIZE)

}

func GetConnect() (*gorm.DB, error) {
	return gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.db"),
	))
}

func Close() {
	err := db.Close()
	if err != nil {
		logs.Logger.Error("gorm.Open close:" + err.Error())
	}
}
