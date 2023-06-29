package goods

import (
	"ET-order-mini-program/database/models"
	"ET-order-mini-program/pkg/middlewares/errorHandling"
	"ET-order-mini-program/pkg/types"
	"ET-order-mini-program/pkg/utils/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandler(r types.Router, db *gorm.DB) {
	goodsRouter := r.Group("/api/goods")
	goodsService := NewGoodsService(db)

	// 创建商品
	goodsRouter.POST("/create", func(c *gin.Context) {
		var dto CreateGoodsDto
		// 获取请求体
		if err := c.ShouldBindJSON(&dto); err != nil {
			panic(errorHandling.NewBadRequestError(err.Error()))
		}
		// 创建实体，写入数据库
		goods, _ := goodsService.Creategoods(&dto)
		// 返回响应
		c.JSON(http.StatusOK,
			models.NewSuccessModel(http.StatusOK, *goods))
	})

	// 查询所有商品
	goodsRouter.GET("/list", func(c *gin.Context) {
		allGoods, err := goodsService.GetAllGoods()
		if err != nil {
			panic(errorHandling.NewNotFoundError(err.Error()))
		}
		c.JSON(http.StatusOK,
			models.NewSuccessModel(http.StatusOK, allGoods))
	})

	// 查询单个商品
	goodsRouter.GET("/detail/:id", func(c *gin.Context) {
		// 获取 param 参数
		id := api.GetIdParamFromCtx(c)
		goods, err := goodsService.GetgoodsById(id)
		if err != nil {
			c.JSON(http.StatusNotFound,
				errorHandling.NewNotFoundError(err.Error()))
			return
		}
		c.JSON(http.StatusOK,
			models.NewSuccessModel(http.StatusOK, *goods))
	})

	// 更新商品
	goodsRouter.PUT("/:id", func(c *gin.Context) {
		// 获取 param 参数
		id := api.GetIdParamFromCtx(c)
		// 获取 body
		var dto UpdateGoodsDto
		if err := c.ShouldBindJSON(&dto); err != nil {
			panic(errorHandling.NewBadRequestError(err.Error()))
		}
		// 判断商品是否存在
		goods := dto.CreateGoods()
		if err := db.First(&goods, id).Error; err != nil {
			panic(errorHandling.NewNotFoundError(err.Error()))
		}
		// 更新记录
		db.Save(&goods)
		c.JSON(http.StatusOK,
			models.NewSuccessModel(http.StatusOK, *goods))
	})

	// 删除商品
	goodsRouter.DELETE("/:id", func(c *gin.Context) {
		// 获取 param 参数
		id := api.GetIdParamFromCtx(c)
		// 查询第一条记录
		var goods models.Goods
		if err := db.First(&goods, id).Error; err != nil {
			panic(errorHandling.NewNotFoundError(err.Error()))
		}
		// 删除记录
		db.Delete(&goods)
		c.JSON(http.StatusOK,
			models.NewSuccessModel(http.StatusOK, goods))
	})
}
