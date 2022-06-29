package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {

	//数据库查询
	return mysql.GetCommunityAll()

}

//GetCommunityDetail 获取社区详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {

	return mysql.GetCommunityDetailById(id)

}
