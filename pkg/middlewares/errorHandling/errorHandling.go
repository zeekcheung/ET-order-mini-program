package errorHandling

import (
	"ET-order-mini-program/database/models"
	"ET-order-mini-program/pkg/middlewares/logger"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一出错处理中间件
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// 捕获错误并输出日志
				logger := logger.GetLoggerFromCtx(c)
				logger.Error(fmt.Sprintf("Panic: %v", r))

				// 根据错误类型返回不同的错误响应
				switch err := r.(type) {
				case string:
					c.JSON(http.StatusInternalServerError,
						models.NewFailModel(http.StatusInternalServerError,
							err),
					)
				case *HttpError:
					c.JSON(err.Code,
						models.NewFailModel(err.Code, err.Message))
				}

				// 终止请求处理
				c.Abort()
			}
		}()

		// 处理请求
		c.Next()
	}
}

// 自定义 HTTP 错误类型
type HttpError struct {
	Code    int
	Message string
}

func (e HttpError) Error() string {
	return e.Message
}

// 创建 HTTP 错误
func NewHTTPError(code int, message string) *HttpError {
	return &HttpError{Code: code, Message: message}
}

// 400 错误
func NewBadRequestError(message string) error {
	return NewHTTPError(http.StatusBadRequest, message)
}

// 401 错误
func NewUnauthorizedError(message string) error {
	return NewHTTPError(http.StatusUnauthorized, message)
}

// 403 错误
func NewForbiddenError(message string) error {
	return NewHTTPError(http.StatusForbidden, message)
}

// 404 错误
func NewNotFoundError(message string) error {
	return NewHTTPError(http.StatusNotFound, message)
}

// 500 错误
func NewInternalServerError(message string) error {
	return NewHTTPError(http.StatusInternalServerError, message)
}
