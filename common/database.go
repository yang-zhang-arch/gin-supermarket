package common

import (
	"WebFull/model"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	driver := viper.GetString("datasource.driver")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	user := viper.GetString("datasource.user")
	charset := viper.GetString("datasource.charset")
	password := viper.GetString("datasource.password")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		user,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(driver, args)

	if err != nil {
		panic("failed to connect to db, err: " + err.Error())
	}
	db.AutoMigrate(&model.User{})

	DB = db

	return db

}

func GetDB() *gorm.DB {
	return DB
}
