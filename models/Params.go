package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

//ParamLogin 登入请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostVoteDate struct {
	//从请求中获取userid
	PostID    string `json:"post_id" binding:"required"`               //帖子id
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1" ` //赞成1 反对-1 取消投票 0
}

//获取帖子列表参数
type ParamPostList struct {
	Page        int64  `from:"page"`
	Size        int64  `from:"size"`
	Order       string `from:"order"`
	CommunityID int64  `json:"community_id" from:"community_id"` //可以为空
}

////获取帖子列表参数
//type ParamCommunityPostList struct {
//	*ParamPostList
//
//}
