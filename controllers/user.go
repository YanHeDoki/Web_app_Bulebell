package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/logic"
	"web_app/models"
	"web_app/utils"
)

func SignUpHandler(c *gin.Context) {

	//1.参数校验
	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(p); err != nil {

		zap.L().Error("ShouldBindJSON err:", zap.Error(err))

		//参数错误直接返回
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.Trans)))

		return
	}
	//2.参数处理

	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("ShouldBindJSON err:")
	//	//参数错误直接返回
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}

	//3.业务处理
	if err := logic.SignUp(p); err != nil {

		zap.L().Error("signup err:", zap.Error(err))
		if errors.Is(err, mysql.ErrUserExist) {
			ResponseError(c, CodeUserExit)
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	//4.返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {

	//请求参数处理以及校验
	p := &models.ParamLogin{}
	if err := c.ShouldBindJSON(p); err != nil {

		zap.L().Error("login with ShouldBindJSON err:", zap.Error(err))

		//参数错误直接返回
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.Trans)))

		return
	}

	//业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {

		zap.L().Error("login.login failed", zap.Error(err))

		if errors.Is(err, mysql.ErrUserNotExist) {
			ResponseError(c, CodeUserNotExit)
		}
		ResponseError(c, CodeInvalidParam)
		return
	}

	//返回响应
	ResponseSuccess(c, gin.H{
		"user_name": user.Username,
		"user_id":   fmt.Sprintf("%d", user.UserId),
		"token":     user.Token,
	})

}
