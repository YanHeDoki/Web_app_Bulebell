package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

//GetCommunityAll 获取社区标签所有标签
func GetCommunityAll() ([]*models.Community, error) {

	sqlstr := "select community_id,community_name from community"
	communities := []*models.Community{}
	err := db.Select(&communities, sqlstr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("this is no community in db")
			return nil, nil
		}
		return nil, err
	}
	return communities, err
}
