package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Response 统一响应结构体
type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// 分页元信息结构体
type PaginationMeta struct {
	PageNum    int `json:"pageNum"`
	PageSize   int `json:"pageSize"`
	Total      int `json:"total"`
	TotalPages int `josn:"totalPages"`
}

// 分页响应结构体
type PaginatedResult[T any] struct {
	Data       []T            `json:"data"`
	Pagination PaginationMeta `json:"pagination"`
}

// 业务状态码到 HTTP 状态码的映射
func codeToHTTPStaus(bizCode int) int {
	switch {
	case bizCode >= 1000 && bizCode < 2000:
		return http.StatusBadRequest
	case bizCode >= 2000 && bizCode < 3000:
		return http.StatusUnauthorized
	case bizCode >= 3000 && bizCode < 4000:
		return http.StatusForbidden
	case bizCode >= 4000 && bizCode < 5000:
		return http.StatusNotFound
	case bizCode >= 5000:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}

func result(c *gin.Context, httpCode, code int, message string, data any) {
	c.JSON(httpCode, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Success(c *gin.Context) {
	result(c, http.StatusOK, CodeSuccess, "成功", nil)
}

func SuccessWithData(c *gin.Context, data any) {
	result(c, http.StatusOK, CodeSuccess, "成功", data)
}

func Error(c *gin.Context, err error) {
	bizErr, ok := err.(*BusinessError)
	if !ok {
		bizErr = ErrServerError
	}
	httpCode := codeToHTTPStaus(bizErr.Code)
	result(c, httpCode, bizErr.Code, bizErr.Message, nil)
}

func ErrorWithData(c *gin.Context, err error, data any) {
	bizErr, ok := err.(*BusinessError)
	if !ok {
		bizErr = ErrServerError
	}
	httpCode := codeToHTTPStaus(bizErr.Code)
	result(c, httpCode, bizErr.Code, bizErr.Message, data)
}

func ValidationError(c *gin.Context, err error) {
	// 类型断言为 ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		// 只取第一个错误返回，使前端处理更简单
		if len(validationErrors) > 0 {
			e := validationErrors[0] // 取第一个错误
			var message string

			// 根据不同的验证规则返回不同的错误信息
			switch e.Tag() {
			case "required":
				message = fmt.Sprintf("%s 是必填项", e.Field())
			case "email":
				message = "请输入有效的邮箱地址"
			case "min":
				if e.Kind().String() == "string" {
					message = fmt.Sprintf("%s 长度不能少于 %s 个字符", e.Field(), e.Param())
				} else {
					message = fmt.Sprintf("%s 不能小于 %s", e.Field(), e.Param())
				}
			case "max":
				if e.Kind().String() == "string" {
					message = fmt.Sprintf("%s 长度不能超过 %s 个字符", e.Field(), e.Param())
				} else {
					message = fmt.Sprintf("%s 不能大于 %s", e.Field(), e.Param())
				}
			case "gte":
				if e.Kind().String() == "int" {
					message = fmt.Sprintf("%s 必须大于或等于 %s", e.Field(), e.Param())
				} else {
					message = fmt.Sprintf("%s 长度必须大于或等于 %s", e.Field(), e.Param())
				}
			case "lte":
				if e.Kind().String() == "int" {
					message = fmt.Sprintf("%s 必须小于或等于 %s", e.Field(), e.Param())
				} else {
					message = fmt.Sprintf("%s 长度必须小于或等于 %s", e.Field(), e.Param())
				}
			case "eqfield":
				message = fmt.Sprintf("%s 必须与 %s 相同", e.Field(), e.Param())
			case "nefield":
				message = fmt.Sprintf("%s 不能与 %s 相同", e.Field(), e.Param())
			case "oneof":
				message = fmt.Sprintf("%s 必须是以下值之一: %s", e.Field(), strings.Replace(e.Param(), " ", ", ", -1))
			default:
				message = fmt.Sprintf("%s 验证失败: %s", e.Field(), e.Tag())
			}

			// 返回单一错误消息
			c.JSON(http.StatusBadRequest, Response{
				Code:    1000,
				Message: message,
				Data:    nil,
			})
			return
		}
	}
	// 如果不是验证错误，返回通用错误
	Error(c, ErrInvalidParams)
}
