package mysql

import "errors"

var (
	ErrUserExist    = errors.New("用户已经存在")
	ErrUserNotExist = errors.New("用户不存在")
	ErrInvalidParam = errors.New("密码或者用户名错误")
	ErrInvalidiID   = errors.New("无效的id")
)
