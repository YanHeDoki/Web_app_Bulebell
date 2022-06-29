package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app/logic"
	"web_app/models"
	"web_app/utils"
)

//投票

func PostVoteController(c *gin.Context) {

	//参数的校验
	p := &models.PostVoteDate{}
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errdata := utils.RemoveTopStruct(errs.Translate(utils.Trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errdata)
		return
	}
	userid, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	err = logic.VoteForPost(userid, p)
	if err != nil {
		zap.L().Error("logic.voteforpost err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}
