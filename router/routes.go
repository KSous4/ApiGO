package router

import (
	UserHandler "github.com/Txug0/ApiGO/handlers/userhandlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	v1 := router.Group("/api/v1/")

	v1.GET("/User", UserHandler.ShowUser)
	v1.POST("/User", UserHandler.CreateUser)
	v1.PUT("/User", UserHandler.UpdateUser)
	v1.DELETE("/User", UserHandler.DeleteUser)
	v1.GET("/Users", UserHandler.ListUsers)

}
