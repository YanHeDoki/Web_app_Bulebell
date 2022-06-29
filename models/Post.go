package models

import "time"

type Post struct {
	ID          int64     `json:"id,string" db:"post_id"`
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"Status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

//获取帖子详情的结构体
type ApiPostDetail struct {
	AuthorName string `json:"author_name"`
	VotesNumb  int64  `json:"votes_numb"`
	*Post
	*CommunityDetail `json:"community_detail"`
}
