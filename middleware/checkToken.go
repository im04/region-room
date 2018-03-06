package middleware

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"log"
	"net/http"
	"region-room/model/error"
)

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("middleware")
		t := time.Now()
		// before request
		token := ctx.GetHeader("token")
		fmt.Println("token:" + token, token == "")
		if token == "" {
			ctx.Abort()
			ctx.JSON(http.StatusOK, gin.H{
				"code": error.ErrorCode.LoginError,
				"msg": "请登录",
			})
		} else {
			ctx.Next()
		}
		// after request
		latency := time.Since(t)
		log.Print(latency)
	}
}