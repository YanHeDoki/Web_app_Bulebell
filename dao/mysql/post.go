package mysql

import (
	"github.com/jmoiron/sqlx"
	"strings"
	"web_app/models"
)

func CreatePost(p *models.Post) error {

	sqlstr := "insert into post(post_id,title,content,author_id,community_id)values(?,?,?,?,?)"
	_, err := db.Exec(sqlstr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	if err != nil {
		return err
	}
	return nil

}

func GetPostById(pid int64) (*models.Post, error) {

	sqlstr := "select post_id,title,community_id,content,author_id,create_time from post where post_id=?"
	p := &models.Post{}
	err := db.Get(p, sqlstr, pid)
	return p, err
}

func GetPostList(page, size int64) (data []*models.Post, err error) {

	sqlstr := "select post_id,title,community_id,content,author_id,create_time from post ORDER BY create_tiem desc limit ?,?"
	data = make([]*models.Post, 0, 2)
	err = db.Select(&data, sqlstr, (page-1)*size, size)

	if err != nil {
		return nil, err
	}
	return
}

func GetPostByIDs(ids []string) (postlist []*models.Post, err error) {

	sqlstr := "select post_id,title,community_id,content,author_id,create_time from post where post_id in (?) order by FIND_IN_SET(post_id,?) "
	query, args, err := sqlx.In(sqlstr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	postlist = make([]*models.Post, 0, 10)
	query = db.Rebind(query)
	err = db.Select(&postlist, query, args...)
	if err != nil {
		return nil, err
	}
	return
}
