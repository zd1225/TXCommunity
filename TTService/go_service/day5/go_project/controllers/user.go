package main

/*
go.mod需要在项目的根目录，否则gopls无法将GET方法自动补全
*/
import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	r := gin.Default()
	r.GET("/getUser", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "success"})
	})

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, Response{
			Code:    200,
			Message: "pong",
		})
	})

	r.Run()
}
