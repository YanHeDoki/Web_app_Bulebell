package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExit
	CodeUserNotExit
	CodeInvalidPassword
	CodeServerBusy

	CodeNeedAuth
	CodeInvalidToken
	CodeNeedLogin
)

var CodemsgMap = map[ResCode]string{
	CodeSuccess:         "请求成功",
	CodeInvalidParam:    "请求参数有误",
	CodeUserExit:        "用户已存在",
	CodeUserNotExit:     "用户不存在",
	CodeInvalidPassword: "用户名或者密码无效",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidToken:    "无效的token",
}

func (rc ResCode) Msg() string {

	msg, ok := CodemsgMap[rc]
	if !ok {
		msg = CodemsgMap[CodeServerBusy]
	}
	return msg
}
