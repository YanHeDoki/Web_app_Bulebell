package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
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
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": utils.RemoveTopStruct(errs.Translate(utils.Trans)),
		})
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

		c.JSON(http.StatusOK, gin.H{
			"msg": "用户注册失败" + err.Error(),
		})
		return
	}

	//4.返回响应
	c.JSONP(http.StatusOK, "ok")
}
