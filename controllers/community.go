package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/logic"
)

//社区相关的接口

func CommunityHandler(c *gin.Context) {

	//获取数据库中的社区的列表（Community_id,Community_name）
	Clist, err := logic.GetCommunityList()

	if err != nil {
		zap.L().Error("logic GetCommunity err:", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不对外暴露服务器问题
		return
	}
	ResponseSuccess(c, Clist)
}
