package user

import (
	"ET-order-mini-program/configs"
	"encoding/json"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// 向微信小程序官方接口发送登录请求
func LoginWechat(code string) (*WechatLoginRespDto, error) {
	client := resty.New()
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"appid":      "小程序 appId",
			"secret":     "小程序 appSecret",
			"js_code":    code,
			"grant_type": "authorization_code",
		}).
		Get("https://api.weixin.qq.com/sns/jscode2session")
	if err != nil {
		return nil, err
	}
	var body WechatLoginRespDto
	if err := json.Unmarshal(resp.Body(), &body); err != nil {
		return nil, err
	}
	return &body, nil
}

// 生成 JWT
func GenerateJWT(openid string, expireAfter time.Duration) (signedToken string, err error) {
	// 设置 token 过期时间
	expireAt := time.Now().Add(expireAfter).Unix()

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": openid,
		"exp":    expireAt,
	})

	// 签名 token
	config, err := configs.LoadConfig(gin.Mode())
	if err != nil {
		return "", err
	}
	secret := config.Secret.TokenKey
	signedToken, err = token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
