package user

import "ET-order-mini-program/database/models"

type LoginDto struct {
	code string
}

type WechatLoginRespDto struct {
	Openid      string
	Session_key string
	Unionid     string
	Errcode     int
	Errmsg      string
}

type UserInfo struct {
	Openid string
}

type LoginRespData struct {
	UserInfo UserInfo
	Token    string
}

type LoginResponse models.ResponseModel[LoginRespData]
