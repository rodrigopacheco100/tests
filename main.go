// @title           API de Clientes
// @version         1.0
// @description     API para gerenciamento de clientes
// @termsOfService  http://swagger.io/terms/

// @contact.name   Suporte
// @contact.email  suporte@example.com

// @host      localhost:3000
// @BasePath  /

package main

import (
	"fmt"
	"tests/infrastructure/env"
	"tests/interfaces"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "tests/docs"
)

func main() {
	app := gin.Default()
	interfaces.SetupRoutes(app)

	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)

	app.GET("/swagger/*any", func(ctx *gin.Context) {
		if ctx.Param("any") == "/" {
			ctx.Request.RequestURI = "/swagger/index.html"
		}
		swaggerHandler(ctx)
	})

	fmt.Println("Servidor rodando em http://localhost:3000")
	app.Run(":3000")
}

func init() {
	env.LoadEnvs()
}
