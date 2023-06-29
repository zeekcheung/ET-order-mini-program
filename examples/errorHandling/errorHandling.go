package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 注册中间件
	r.Use(errorHandlerMiddleware)

	// 测试路由
	r.GET("/400", func(c *gin.Context) {
		panic(newBadRequestError("bad request"))
		// c.String(http.StatusOK, "Hello, world!")
	})

	r.GET("/401", func(c *gin.Context) {
		panic(newUnauthorizedError("unauthorizaed"))
		// c.String(http.StatusOK, "Hello, world!")
	})

	r.GET("/403", func(c *gin.Context) {
		panic(newForbiddenError("fobidden"))
		// c.String(http.StatusOK, "Hello, world!")
	})

	r.GET("/500", func(c *gin.Context) {
		panic(newInternalServerError())
		// c.String(http.StatusOK, "Hello, world!")
	})

	r.Run(":8888")
}

// 统一出错处理中间件
func errorHandlerMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			// 捕获错误并输出日志
			fmt.Printf("Panic: %v\n", err)

			// 根据错误类型返回不同的错误响应
			switch err.(type) {
			case string:
				c.JSON(http.StatusBadRequest, gin.H{"error": err})
			case *httpError:
				httpErr := err.(*httpError)
				c.JSON(httpErr.status, gin.H{"error": httpErr.message})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			}

			// 终止请求处理
			c.Abort()
		}
	}()

	// 处理请求
	c.Next()
}

// 自定义 HTTP 错误类型
type httpError struct {
	status  int
	message string
}

func (e httpError) Error() string {
	return e.message
}

// 创建 HTTP 错误
func newHTTPError(status int, message string) *httpError {
	return &httpError{status: status, message: message}
}

// 400 错误
func newBadRequestError(message string) error {
	return newHTTPError(http.StatusBadRequest, message)
}

// 401 错误
func newUnauthorizedError(message string) error {
	return newHTTPError(http.StatusUnauthorized, message)
}

// 403 错误
func newForbiddenError(message string) error {
	return newHTTPError(http.StatusForbidden, message)
}

// 404 错误
func newNotFoundError(message string) error {
	return newHTTPError(http.StatusNotFound, message)
}

// 500 错误
func newInternalServerError() error {
	return newHTTPError(http.StatusInternalServerError, "Internal Server Error")
}
