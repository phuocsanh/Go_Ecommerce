package initialize

import (
	"fmt"
	"go_ecommerce/global"
)

func Run(){
	LoadConfig()
	fmt.Println("Load configuration mysql", global.Config.Mysql.Username)
	InitLoger()
	InitMysql()
	InitRedis()

	r:=InitRouter()
	r.Run(":8002")
}