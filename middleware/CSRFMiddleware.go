package middleware

import "github.com/gin-gonic/gin"

func CSRFMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	}

}
