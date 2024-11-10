package main

import (
	_ "go_ecommerce/cmd/swag/docs"
	"go_ecommerce/internal/initialize"

	swaggerFiles "github.com/swaggo/files"     //
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Api documentation ecommerce_sq
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  https://github.com/phuocsanh/Go_Ecommerce-go

// @contact.name   API Support
// @contact.url    https://github.com/phuocsanh/Go_Ecommerce-go
// @contact.email  phuocsanhtps@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /api/v1

func main() {
	r := initialize.Run()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8002")
}