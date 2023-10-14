package UserHandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "deu certo essa merda",
	})
}
