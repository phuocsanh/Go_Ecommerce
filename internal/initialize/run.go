package initialize

import (
	"fmt"
	"go_ecommerce/global"

	"go.uber.org/zap"
)

func Run(){
	LoadConfig()
	fmt.Println("Load configuration mysql", global.Config.Mysql.Username)
	InitLogger()

	global.Logger.Debug("config log ok", zap.String("ok", "success"))
	InitMysql()
	InitMysqlC()
	InitServiceInterface()
	InitRedis()

	r:=InitRouter()
	r.Run(":8002")
}