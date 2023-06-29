package shop

import (
	"ET-order-mini-program/database/models"
	"ET-order-mini-program/pkg/middlewares/errorHandling"
	"ET-order-mini-program/pkg/types"
	"ET-order-mini-program/pkg/utils/conv"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandler(r types.Router, db *gorm.DB) {
	shopRouter := r.Group("/api/shop")
	shopService := NewShopService(db)

	// TODO: 创建店铺
	// shopRouter.POST("/create", middlewares.AuthMiddleware(), func(c *gin.Context) {
	shopRouter.POST("/create", func(c *gin.Context) {
		var dto CreateShopDto
		// 获取请求体
		if err := c.ShouldBindJSON(&dto); err != nil {
			panic(errorHandling.NewBadRequestError(err.Error()))
		}
		// 创建实体，写入数据库
		shop, _ := shopService.CreateShop(&dto)
		// 返回响应
		c.JSON(http.StatusOK,
			models.NewSuccessModel(http.StatusOK, *shop))
	})

	// TODO: 查询所有店铺
	shopRouter.GET("/list", func(c *gin.Context) {
		allShops, err := shopService.GetAllShops()
		if err != nil {
			panic(errorHandling.NewNotFoundError(err.Error()))
		}
		c.JSON(http.StatusOK,
			models.NewSuccessModel(http.StatusOK, allShops))
	})

	// TODO: 查询店铺详情
	shopRouter.GET("/detail/:id", func(c *gin.Context) {
		// 获取 param 参数
		id, ok := c.Params.Get("id")
		if !ok {
			panic(errorHandling.NewBadRequestError("can not get the param: id"))
		}
		shop, err := shopService.GetShopById(conv.ConvStrToUint(id))
		if err != nil {
			c.JSON(http.StatusNotFound,
				errorHandling.NewNotFoundError(err.Error()))
			return
		}
		c.JSON(http.StatusOK,
			models.NewSuccessModel(http.StatusOK, *shop))
	})

	// // 更新店铺
	// shopRouter.POST("/update", func(c *gin.Context) {
	// 	// 获取 body
	// 	var dto UpdateGoodsDto
	// 	if err := c.ShouldBindJSON(&dto); err != nil {
	// 		c.AbortWithError(http.StatusBadRequest, err)
	// 		return
	// 	}
	// 	// 判断商品是否存在
	// 	goods := dto.CreateGoods()
	// 	if err := db.First(&goods, id).Error; err != nil {
	// 		c.AbortWithError(http.StatusNotFound,
	// 			errors.New("can not get the param: id"))
	// 		return
	// 	}
	// 	// 更新记录
	// 	db.Save(&goods)
	// 	c.JSON(http.StatusOK,
	// 		models.ResponseModel[models.Goods]{Data: *goods})
	// })

	// // 删除店铺
	// shopRouter.DELETE("/:id", func(c *gin.Context) {
	// 	// 获取 param 参数
	// 	id, ok := c.Params.Get("id")
	// 	if !ok {
	// 		c.AbortWithError(http.StatusBadRequest,
	// 			errors.New("can not get the param: id"))
	// 		return
	// 	}
	// 	var goods models.Goods
	// 	// 查询第一条记录
	// 	if err := db.First(&goods, id).Error; err != nil {
	// 		c.AbortWithError(http.StatusNotFound, err)
	// 		return
	// 	}
	// 	// 删除记录
	// 	db.Delete(&goods)
	// 	c.JSON(http.StatusOK, models.ResponseModel[models.Goods]{Data: goods})
	// })
}
