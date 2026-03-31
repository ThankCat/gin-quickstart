package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const learnGinDocsURL = "https://gin-gonic.com/zh-cn/docs/routing/api-design/"

type Response struct {
	Success bool       `json:"success"`
	Data    any        `json:"data,omitempty"`
	Error   *ErrorInfo `json:"error,omitempty"`
}

type ErrorInfo struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"pre_page,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}

func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func Fail(c *gin.Context, status int, code string, message string) {
	c.JSON(status, Response{
		Success: false,
		Error:   &ErrorInfo{Code: code, Message: message},
	})
}

func main() {
	router := gin.Default()

	router.GET("/api/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "0" {
			Fail(ctx, http.StatusNotFound, "USER_NOT_FOUND", "no user with that ID")
		} else {
			OK(ctx, gin.H{"id": id, "name": "Alice"})
		}
	})

	router.Run(":8080") // listens on 0.0.0.0:8080 by default
}
