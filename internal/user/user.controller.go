package user

import (
	"ET-order-mini-program/pkg/types"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandler(r types.Router, db *gorm.DB) {
	userRouter := r.Group("/api/user")

	// TODO: 登录
	userRouter.POST("/login", func(c *gin.Context) {
		var dto LoginDto
		// 获取请求体
		if err := c.ShouldBindJSON(&dto); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		// 向微信小程序官方接口发送登录请求
		resp, err := LoginWechat(dto.code)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// 生成 token
		token, err := GenerateJWT(resp.Openid, 24*time.Hour)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		// 返回响应
		c.JSON(http.StatusOK, LoginResponse{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
			Data: LoginRespData{
				UserInfo: UserInfo{
					Openid: resp.Openid,
				},
				Token: token,
			},
		})
	})
}
