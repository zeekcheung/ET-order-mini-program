package api

import (
	"ET-order-mini-program/pkg/middlewares/errorHandling"

	"github.com/gin-gonic/gin"
)

func GetIdParamFromCtx(c *gin.Context) string {
	// 获取 param 参数
	id, ok := c.Params.Get("id")
	if !ok {
		panic(errorHandling.NewBadRequestError("can not get the param: id"))
	}
	return id
}
