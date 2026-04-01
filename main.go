package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const learnGinDocsURL = "https://gin-gonic.com/zh-cn/docs/routing/api-design/"

type AppError struct {
	Status  int    `json:"-"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrNotFound     = &AppError{Status: 404, Code: "NOT_FOUND", Message: "resource not found"}
	ErrUnauthorized = &AppError{Status: 401, Code: "UNAUTHORIZED", Message: "authentication required"}
	ErrBadRequest   = &AppError{Status: 400, Code: "BAD_REQUEST", Message: "invalid request"}
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) == 0 {
			return
		}

		err := ctx.Errors.Last().Err
		var appErr *AppError
		if errors.As(err, &appErr) {
			ctx.JSON(appErr.Status, gin.H{
				"success": false,
				"error":   gin.H{"code": appErr.Code, "message": appErr.Message},
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   gin.H{"code": "INTERNAL", "message": "an unexpected error occurred"},
			})
		}
	}
}

func main() {
	router := gin.Default()

	router.GET("/api/items/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		if id == "0" {
			ctx.Error(ErrNotFound)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"id": id}})
	})

	router.Run(":8080") // listens on 0.0.0.0:8080 by default
}
