package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingGet(c *gin.Context){
	c.JSON(http.StatusOK,map[string] string{
		"Hello" : "Found me",
	})
}
