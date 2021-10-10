package router

import (
	"github.com/Hidemasa-Kajita/go_api_sample/controller"
	"github.com/gin-gonic/gin"
)

func Run() {
	engine := gin.Default()

	users := engine.Group("/users")
	{
		users.GET("", controller.Index)
		users.GET(":id", controller.Show)
		users.POST("", controller.Store)
		users.PUT(":id", controller.Update)
		users.DELETE(":id", controller.Delete)
	}

	engine.Run(":8080")
}
