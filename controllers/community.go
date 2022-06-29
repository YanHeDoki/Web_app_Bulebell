package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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

//CommunityDetailHandler 根据id获取社区分类详情
func CommunityDetailHandler(c *gin.Context) {

	//gin 获取路由参数
	CommunityId := c.Param("id")
	Id, err := strconv.ParseInt(CommunityId, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	//获取数据库中的社区的列表（Community_id,Community_name）
	communityDetail, err := logic.GetCommunityDetail(Id)

	if err != nil {
		zap.L().Error("logic GetCommunity err:", zap.Error(err))
		ResponseError(c, CodeServerBusy) //不对外暴露服务器问题
		return
	}
	ResponseSuccess(c, communityDetail)
}
