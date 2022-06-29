package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/models"
	"web_app/snowflake"
)

func CreatePost(p *models.Post) error {

	//生成帖子id
	p.ID = snowflake.GenID()

	//保存到数据库
	err := mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	if err != nil {
		return err
	}
	return err
}

//GetPostByid 根据帖子id 获取详情
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {

	//查询并且处理接口的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		return data, err
	}

	var name string
	name, err = mysql.GetUserById(post.AuthorID)
	if err != nil {
		return data, err
	}
	communityDetail, err := mysql.GetCommunityDetailById(post.CommunityID)
	if err != nil {
		return data, err
	}
	data = &models.ApiPostDetail{
		AuthorName:      name,
		Post:            post,
		CommunityDetail: communityDetail,
	}

	return data, nil
}

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {

	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		name, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			return data, err
		}
		communityDetail, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			return data, err
		}
		Detail := &models.ApiPostDetail{
			AuthorName:      name,
			Post:            post,
			CommunityDetail: communityDetail,
		}

		data = append(data, Detail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {

	//去redis中取值
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) success but return 0 ids")
		return
	}

	//提前去redis中查询每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	//根据id去查询数据库帖子详情信息
	posts, err := mysql.GetPostByIDs(ids)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		name, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			return data, err
		}
		communityDetail, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			return data, err
		}
		Detail := &models.ApiPostDetail{
			AuthorName:      name,
			VotesNumb:       voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		}

		data = append(data, Detail)
	}
	return

}

func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {

	//去redis中取值
	ids, err := redis.GetCommunityPostIDsInOrder(p)
	if err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) success but return 0 ids")
		return
	}

	//提前去redis中查询每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	//根据id去查询数据库帖子详情信息
	posts, err := mysql.GetPostByIDs(ids)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for idx, post := range posts {
		name, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			return data, err
		}
		communityDetail, err := mysql.GetCommunityDetailById(post.CommunityID)
		if err != nil {
			return data, err
		}
		Detail := &models.ApiPostDetail{
			AuthorName:      name,
			VotesNumb:       voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		}

		data = append(data, Detail)
	}
	return
}

func GetCommunityPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {

	if p.CommunityID == 0 {
		data, err = GetPostList2(p)
	} else {
		data, err = GetCommunityPostListNew(p)
	}
	if err != nil {
		return nil, err
	}
	return data, err

}
