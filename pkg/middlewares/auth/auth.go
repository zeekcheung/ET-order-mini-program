package auth

import (
	"ET-order-mini-program/configs"
	"ET-order-mini-program/pkg/middlewares/errorHandling"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TODO: JWT 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			panic(errorHandling.NewUnauthorizedError("missing token"))
		}

		// 解析 token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 检查签名算法是否正确
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			config, err := configs.LoadConfig(gin.Mode())
			if err != nil {
				return "", err
			}
			secret := config.Secret.TokenKey

			// 返回密钥
			return []byte(secret), nil
		})

		if err != nil {
			panic(errorHandling.NewUnauthorizedError(err.Error()))
		}

		// 检查 token 是否有效
		if !token.Valid {
			panic(errorHandling.NewUnauthorizedError("invalid token"))
		}

		// 将 userId 保存到上下文中
		userId, ok := token.Claims.(jwt.MapClaims)["userId"].(float64)
		if !ok {
			panic(errorHandling.NewUnauthorizedError("invalid token"))
		}
		c.Set("userId", int(userId))

		c.Next()
	}
}
