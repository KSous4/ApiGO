package router

import (
	UserHandler "github.com/Txug0/ApiGO/handlers/userhandlers"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	userRoutesGroup := router.Group("/api/v1/user/")

	userRoutesGroup.POST("/createUser", UserHandler.CreateUser)

}
