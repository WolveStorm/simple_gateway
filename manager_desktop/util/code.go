package util

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServeBusy
	CodeEmptyRequestHeader
	CodeInvalidAuth
	CodeNeedLogin
	CodeRepeatLogin
	CodeNameExist
)

var codeMap = map[ResCode]string{
	CodeSuccess:            "请求成功",
	CodeInvalidParam:       "请求参数异常",
	CodeUserExist:          "用户已存在",
	CodeUserNotExist:       "用户不存在",
	CodeInvalidPassword:    "密码错误",
	CodeServeBusy:          "服务器繁忙",
	CodeEmptyRequestHeader: "需要登录",
	CodeInvalidAuth:        "无效的token",
	CodeNeedLogin:          "需要登录",
	CodeRepeatLogin:        "不要重复登陆",
	CodeNameExist:          "服务名已经存在",
}

func Msg(code ResCode) string {
	s, ok := codeMap[code]
	if !ok {
		return codeMap[CodeServeBusy]
	} else {
		return s
	}
}
