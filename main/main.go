package main

import (
	"WebFull/common"
	"WebFull/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()
	app := gin.Default()

	// 添加路由
	app = routes.UserRoute(app)
	port := viper.GetString("server.port")
	app.Run(":" + port)

}

func InitConfig() {
	workDir, _ := os.Getwd()
	fmt.Println(workDir)
	viper.SetConfigType("yml")
	viper.SetConfigName("application")
	viper.AddConfigPath(workDir + "\\config")
	err := viper.ReadInConfig()

	if err != nil {
		panic("配置文件加载错误")
	}
}
