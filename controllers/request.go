package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")
var ContextUserIDKey = "userid"

//GetCurrentUser 通过这个函数快速获取用户id
func GetCurrentUser(c *gin.Context) (int64, error) {

	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err := ErrorUserNotLogin
		return 0, err
	}
	userID, ok := uid.(int64)
	if !ok {
		err := ErrorUserNotLogin
		return 0, err
	}
	return userID, nil
}
