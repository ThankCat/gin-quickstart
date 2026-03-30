package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const learnGinDocsURL = "https://gin-gonic.com/zh-cn/docs/routing/redirects/"

func AuthRequired(ctx *gin.Context) {
	fmt.Println("我经历了一个授权中间件")
	ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
		"msg": "你没有权限",
	})
}

func loginHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"action": "login"})
}

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	v1 := router.Group("/v1")
	v1.POST("/login", loginHandle)

	v2 := router.Group("/v2")
	v2.Use(AuthRequired)
	v2.POST("/login", loginHandle)

	router.Run(":8080") // listens on 0.0.0.0:8080 by default
}
