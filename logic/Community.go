package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
)

func GetCommunityList() ([]*models.Community, error) {

	//数据库查询
	return mysql.GetCommunityAll()

}
