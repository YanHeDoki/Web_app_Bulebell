package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
	"web_app/models"
)

func CreatePostHandler(c *gin.Context) {

	//获取参数以及校验
	p := &models.Post{}
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("createpost err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//从context里面获取用户id
	userid, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userid
	//创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("postCreate err:", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//返回响应
	ResponseSuccess(c, nil)
}

//GetPostDetailHandler 获取帖子详情的
func GetPostDetailHandler(c *gin.Context) {
	//1.获取参数（Url中的帖子的id），根据id去取出帖子数据
	pid := c.Param("id")
	id, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		zap.L().Error("get post detail err: invalid err:", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetPostById(id)
	if err != nil {
		zap.L().Error("get post detail err: invalid err:", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	ResponseSuccess(c, data)
}

func GetPostListHandler(c *gin.Context) {

	//获取数据

	pagestr := c.Query("page")
	sizestr := c.Query("size")

	page, err := strconv.ParseInt(pagestr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseInt(sizestr, 10, 64)
	if err != nil {
		size = 10
	}

	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("get post list  err:", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler2 升级版帖子列表接口
//GetPostListHandler2 根据前端传入的参数动态获取帖子列表
//按照创造时间或者是分数排序
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListHandler2(c *gin.Context) {

	//获取数据
	p := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime,
	}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//从redis中获取id值

	//根据redis的id获取数据库中详细信息

	data, err := logic.GetCommunityPostListNew(p)
	if err != nil {
		zap.L().Error("get post list  err:", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	//返回响应
	ResponseSuccess(c, data)
}

//GetCommunityListHandler 根据社区查询帖子 废弃
func GetCommunityListHandler(c *gin.Context) {

	//获取数据
	p := &models.ParamPostList{}

	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}

	data, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("get post list  err:", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//返回响应
	ResponseSuccess(c, data)

}
