package ginplus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/marspere/goencrypt"
	"github.com/tulanz/pkg/errors"
	"github.com/tulanz/pkg/util"
)

// ResponseAccess ResponseAccess
type ResponseAccess struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}
type ResponsePage struct {
	ResponseAccess
	Pagination *util.Page `json:"page_info,omitempty"`
}

// ResponseList ResponseList
type ResponseList struct {
	List       interface{} `json:"list"`
	Total      int64       `json:"count"` // 统计
	Pagination *Paging     `json:"pagination,omitempty"`
}

// ResponseError ResponseError
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// QueryParams QueryParams
type QueryParams struct {
	PageNum   int // 分页计算
	PageSize  int // 分页条数
	PageIndex int // 当前页
}
type Paging struct {
	PageIndex int `json:"pageIndex" form:"page_index"`
	PageSize  int `json:"pageSize" form:"page_size"`
}

// ValidatorErrorHandle ValidatorErrorHandle
func ValidatorErrorHandle(c *gin.Context, err error) {
	if err, exist := err.(validator.ValidationErrors); exist {
		// logger.Error(err)
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	ResError(c, errors.NewResponse(1000, err.Error(), 200))
}

// GetPaging GetPaging
func GetPaging(c *gin.Context) *Paging {
	var r Paging
	if err := c.ShouldBind(&r); err != nil {
		r.PageIndex = 1
		r.PageSize = 10
	}

	if r.PageIndex < 1 {
		r.PageIndex = 1
	}
	if r.PageSize <= 0 {
		r.PageSize = 20
	}
	return &Paging{
		PageIndex: r.PageIndex,
		PageSize:  r.PageSize,
	}
}

func FotmatPaging(r *Paging) *Paging {
	if r.PageIndex < 1 {
		r.PageIndex = 1
	}
	if r.PageSize <= 0 {
		r.PageSize = 10
	}
	return r
}

func GetPageNum(c *gin.Context) int {
	page := GetPageIndex(c)
	return (page - 1) * GetPageSize(c)
}

// GetPageIndex 获取分页的页索引
func GetPageIndex(c *gin.Context) int {
	defaultVal := 1
	if v := c.Query("page"); v != "" {
		if iv := util.S(v).DefaultInt(defaultVal); iv > 0 {
			return iv
		}
	}
	return defaultVal
}

// GetPageSize 获取分页的页大小
func GetPageSize(c *gin.Context) int {
	defaultVal := 10 //setting.AppSetting.PageSize
	if v := c.Query("limit"); v != "" {
		if iv := util.S(v).DefaultInt(defaultVal); iv > 0 {
			return iv
		} else if iv == -1 {
			return 10000000
		}
	}
	return defaultVal
}

// ResPage 响应分页数据
func ResPage(c *gin.Context, v *ResponseList) {
	ResSuccess(c, v)
}

// ResOK 响应OK
func ResOK(c *gin.Context) {
	result := ResponseAccess{
		Code:    200,
		Message: "操作成功",
	}
	ResJSON(c, http.StatusOK, result)
}

// Ok Ok
func Ok(c *gin.Context, data interface{}) {
	result := ResponseAccess{
		Code:    0,
		Message: "操作成功",
		Data:    data,
	}
	ResJSON(c, http.StatusOK, result)
	// c.AbortWithStatusJSON(http.StatusOK, &result)
}

// Encrypt 加密数据
func Encrypt(c *gin.Context, key string, data interface{}) {
	buf, err := util.JSONMarshal(data)
	if err != nil {
		panic(err)
	}

	cipher := goencrypt.NewDESCipher([]byte(key), []byte(""), goencrypt.CBCMode, goencrypt.PkcsZero, goencrypt.PrintBase64)
	cipherText, _ := cipher.TripleDESEncrypt(buf)
	result := ResponseAccess{
		Code:    0,
		Message: "操作成功",
		Data:    cipherText,
	}
	ResJSON(c, http.StatusOK, result)
}

func Page(c *gin.Context, paging *Paging, total int64, list interface{}) {
	v := ResponseList{
		List:       list,
		Total:      total,
		Pagination: paging,
	}
	Ok(c, &v)
}

// Simple Simple
func Simple(c *gin.Context, err string) {
	data := map[string]interface{}{
		"msg": err,
	}
	ResJSON(c, http.StatusOK, data)
}

// ResSuccess 响应成功
func ResSuccess(c *gin.Context, v interface{}) {
	ResJSON(c, http.StatusOK, v)
}

// ResJSON 响应JSON数据
func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := util.JSONMarshal(v)
	if err != nil {
		panic(err)
	}
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

// ResError 响应错误
func ResError(c *gin.Context, err error, status ...int) {
	var resError *errors.ResponseError
	if err != nil {
		if re, ok := err.(*errors.ResponseError); ok {
			resError = re
		} else {
			resError = errors.UnWrapResponse(errors.Wrap500Response(err))
		}
	} else {
		resError = errors.UnWrapResponse(errors.ErrInternalServer)
	}

	dtoError := ResponseError{
		Code:    resError.Code,
		Message: resError.Message,
	}

	if err := resError.ERR; err != nil {
		dtoError.Error = err.Error()
		if status := resError.Status; status >= 400 && status < 500 {
			// logger.StartSpan(NewContext(c)).Warnf(err.Error())
		} else if status >= 500 {
			// span := logger.StartSpan(NewContext(c))
			// span = span.WithField("stack", fmt.Sprintf("%+v", err))
			// span.Errorf(err.Error())
		}
	}
	ResJSON(c, http.StatusOK, dtoError)
}

// Error Error
func Error(c *gin.Context, code int, err error) {
	resp := &ResponseError{
		Code:    code,
		Message: err.Error(),
	}
	ResJSON(c, http.StatusOK, resp)
}
