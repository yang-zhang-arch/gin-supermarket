package middleware

import (
	"WebFull/common"
	"WebFull/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" || !(strings.HasPrefix(token, "Bearer ")) {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "认证失败",
			})
			ctx.Abort()
			return
		}
		tokenStr := token[7:]
		tok, claim, err := common.ParseToken(tokenStr)
		if err != nil || !tok.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "认证失败",
			})
			ctx.Abort()
			return
		}
		// OPT: 查库验证，生产环境可以跳过此步骤
		userId := claim.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		if user.ID == 0 {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"code":    "422",
				"message": "用户不存在，Token已失效",
			})
			ctx.Abort()
			return
		}
		ctx.Set("user", user)
		ctx.Next()
	}
}
