package errors

import (
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/wonderivan/logger"
)

var (
	trans        ut.Translator
	uni          *ut.UniversalTranslator
	errorHandler *AppErrorHandler
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		uni = ut.New(zh.New())
		trans, _ = uni.GetTranslator("zh")
		// 收集结构体中的comment标签，用于替换英文字段名称
		v.SetTagName("validate")
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return fld.Tag.Get("comment")
		})
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		zh_translations.RegisterDefaultTranslations(v, trans)
		errorHandler = NewErrorHandler(uni, trans, v)
	}
}

// ValidateError 错误处理
func ValidateError() gin.HandlerFunc {
	return func(c *gin.Context) {
		errorHandler.HandleErrors(c)
	}
}

// AppErrorHandler AppErrorHandler
type AppErrorHandler struct {
	uni      *ut.UniversalTranslator
	trans    ut.Translator
	validate *validator.Validate
}

// NewErrorHandler NewErrorHandler
func NewErrorHandler(uni *ut.UniversalTranslator, trans ut.Translator, validate *validator.Validate) *AppErrorHandler {
	return &AppErrorHandler{
		uni:      uni,
		trans:    trans,
		validate: validate,
	}
}

// HandleErrors HandleErrors
func (h *AppErrorHandler) HandleErrors(c *gin.Context) {
	defer func() {
		errorToPrint := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if errorToPrint != nil {
			if errs, ok := errorToPrint.Err.(validator.ValidationErrors); ok {
				trans, _ := h.uni.GetTranslator("zh")
				errors := make(map[string]interface{})
				for _, v := range errs {
					errors[v.StructNamespace()] = v.Translate(trans)
				}
				//打印错误堆栈信息
				logger.Debug("panic: %v\n", errs)
				// debug.PrintStack()
				//封装通用json返回

				v := gin.H{
					"code":      http.StatusBadRequest,
					"message":   errs[0].Translate(trans),
					"errors":    errors,
					"timestamp": time.Now().UnixNano() / 1000 / 1000,
					"error":     http.StatusText(http.StatusBadRequest),
					"path":      c.Request.URL.Path,
				}

				c.JSON(http.StatusOK, v)
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()

			}
		}
	}()
	c.Next()
}
