package errors

import (
	"github.com/pkg/errors"
)

// 定义别名
var (
	Is           = errors.Is
	As           = errors.As
	New          = errors.New
	Wrap         = errors.Wrap
	Wrapf        = errors.Wrapf
	Cause        = errors.Cause
	Unwrap       = errors.Unwrap
	WithStack    = errors.WithStack
	WithMessage  = errors.WithMessage
	WithMessagef = errors.WithMessagef
)

// 定义错误
var (
	ErrBadRequest              = New400Response("请求发生错误")
	ErrInvalidParams           = New400Response("请求参数错误")
	ErrInvalidParent           = New400Response("无效的父级节点")
	ErrNotAllowDeleteWithChild = New400Response("含有子级，不能删除")
	ErrNotAllowDelete          = New400Response("资源不允许删除")
	ErrInvalidUserName         = New400Response("无效的用户名")
	ErrInvalidPassword         = New400Response("无效的密码")
	ErrInvalidUser             = New400Response("无效的用户")
	ErrUserDisable             = New400Response("您的账户已被冻结/停用，如有疑问，请和客服人员联系～")

	ErrInvalidCode     = NewResponse(1002, "无效的验证码", 200)
	ErrNoPerm          = NewResponse(401, "无访问权限", 401)
	ErrInvalidToken    = NewResponse(9999, "令牌失效", 406)
	ErrTimeOutToken    = NewResponse(10000, "令牌已过期", 406)
	ErrNotFound        = NewResponse(404, "资源不存在", 404)
	ErrMethodNotAllow  = NewResponse(405, "方法不被允许", 405)
	ErrTooManyRequests = NewResponse(429, "请求过于频繁", 429)
	ErrInternalServer  = NewResponse(500, "服务器发生错误", 500)

	// Err_DB_QUERY = NewResponse(204, "")
	E_UNKNOW              = NewResponse(-1, "未知错误")
	E_HTTP_PARAMETER      = NewResponse(101, "请求参数错误")
	E_HTTP_URL            = NewResponse(102, "请求地址错误")
	E_HTTP_METHOD         = NewResponse(103, "请求方法错误")
	E_HTTP_IP             = NewResponse(104, "请求IP错误")
	E_HTTP_NewResponse    = NewResponse(105, "请求内部错误")
	E_HTTP_TIMEOUT        = NewResponse(106, "请求超时")
	E_DB_ADD              = NewResponse(201, "数据库新增错误")
	E_DB_DELETE           = NewResponse(202, "数据库删除错误")
	E_DB_UPDATE           = NewResponse(203, "数据库更新错误")
	E_DB_QUERY            = NewResponse(204, "数据库查询错误")
	E_DB_TRANSACTION      = NewResponse(205, "数据库事务失败")
	E_TOKEN_EMPTY         = NewResponse(301, "token未找到")
	E_TOKEN_EXPIRED       = NewResponse(302, "token过期")
	E_TOKEN_INVALID       = NewResponse(303, "token非法")
	E_TOKEN_NOT_MATCH     = NewResponse(304, "token不匹配")
	E_TOKEN_WEIXIN_FAILED = NewResponse(305, "token获取失败")
	E_MSG_SEND            = NewResponse(401, "发送短信验证码失败")
	E_MSG_SEND_MANY       = NewResponse(402, "发送短信验证码失败，超过个人每日接收短信最大数目")
	E_MSG_SEND_FREQ       = NewResponse(403, "发送短信验证码失败，发送过于频繁，请稍等再试")
	E_MSG_SEND_VERIFY     = NewResponse(404, "发送短信验证码失败，输入错误")
	E_MSG_BLACK_LIST      = NewResponse(405, "发送短信验证码失败，手机号码处于黑名单中")
	E_MSG_VERIFY_FAILED   = NewResponse(406, "短信码验证失败")
	E_MSG_EXPIRED         = NewResponse(406, "短信码验证过期")
	E_FILE_EXPORT         = NewResponse(501, "文件导出失败")
	E_FILE_UPLOAD         = NewResponse(502, "文件上传失败")
	E_JSON_MARSHA         = NewResponse(601, "json序列化失败")
	E_JSON_UNMARSHA       = NewResponse(602, "json反序列化失败")
)
