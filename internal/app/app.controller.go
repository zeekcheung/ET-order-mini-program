package app

import (
	"ET-order-mini-program/configs"
	"ET-order-mini-program/internal/goods"
	"ET-order-mini-program/internal/shop"
	"ET-order-mini-program/internal/user"
	"ET-order-mini-program/pkg/middlewares/errorHandling"
	"ET-order-mini-program/pkg/middlewares/logger"
	"ET-order-mini-program/pkg/types"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 初始化路由
func InitRouter(db *gorm.DB) (router types.Router) {
	// 创建路由实例
	engine := gin.Default()
	router = types.Router{Engine: engine}
	// 注册中间件
	router.Use(logger.LoggerMiddleware())
	router.Use(errorHandling.ErrorHandlerMiddleware())
	// 注册路由
	RegisterHandlers(router, db)
	return router
}

// 注册路由
func RegisterHandlers(r types.Router, db *gorm.DB) {
	user.RegisterHandler(r, db)
	shop.RegisterHandler(r, db)
	goods.RegisterHandler(r, db)
}

// 启动服务器
func Serve(r types.Router, config *configs.Config) {
	env := gin.Mode()
	serverConfig := config.Server
	addr := fmt.Sprintf("%s:%s", serverConfig.Host, serverConfig.Port)

	var err error
	if env == gin.DebugMode {
		// 开发环境
		err = r.Run(addr)
	} else if env == gin.ReleaseMode {
		// 生产环境
		err = r.RunTLS(addr, "./key/server.crt", "./key/server.key")
	} else if env == gin.TestMode {
		// 测试环境
		err = r.RunTLS(addr, "./key/server.crt", "./key/server.key")
	}

	if err != nil {
		panic("Failed to start server, error: " + err.Error())
	}
}
