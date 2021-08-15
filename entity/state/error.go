package state

import (
	"fmt"
	"net/http"

	"github.com/go-zelus/zelus/core/config"
)

var (
	status = map[int32]string{}
)

type (
	Error struct {
		code    int32
		message string
	}

	ErrorResponse struct {
		Code    int32  `json:"code"`    // 错误代码
		Message string `json:"message"` // 错误描述
	}
)

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() int32 {
	return e.code
}

func (e *Error) Message(args ...interface{}) string {
	if len(args) == 0 {
		return e.message
	}
	res := make([]interface{}, 0)
	res = append(res, e.message)
	if config.GetBool("app.debug") {
		res = append(res, "-")
		res = append(res, args...)
	}
	return fmt.Sprint(res...)
}

func (e *Error) Response(args ...interface{}) *ErrorResponse {
	return &ErrorResponse{
		Code:    e.Code(),
		Message: e.Message(args...),
	}
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		fallthrough
	case UsernameExisted.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}

func NewError(code int32, message string) *Error {
	if _, ok := status[code]; ok {
		panic(fmt.Sprintf("状态码 %d 已存在", code))
	}
	status[code] = message
	return &Error{
		code:    code,
		message: message,
	}
}

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000, "服务内部错误")
	InvalidParams             = NewError(10001, "传入参数错误")
	NotFound                  = NewError(10002, "未找到对应的资源")
	UnauthorizedAuthNotExist  = NewError(10003, "鉴权失败，未找到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(10005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10007, "请求过多")
	DatabaseError             = NewError(10008, "数据库处理错误")
	RedisError                = NewError(10009, "缓存处理错误")

	IDIllegal       = NewError(20000, "非法的编号")
	UsernameExisted = NewError(20001, "用户名已存在")
	MobileExisted   = NewError(20002, "手机号已存在")
	EmailExisted    = NewError(20003, "邮箱已存在")
	LoginError      = NewError(20004, "账号/密码错误")
	SMSError        = NewError(20005, "短信发送失败")
)
