package redis

import (
	"github.com/go-redis/redis"
	"strconv"
	"time"
	"web_app/models"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {

	//获取从redis获取id
	start := (page - 1) * size
	end := start + size - 1
	//redis 查询
	return rdb.ZRevRange(key, start, end).Result()
}

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {

	//获取从redis获取id
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScoreZSet)
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//redis 查询
	return rdb.ZRevRange(key, start, end).Result()

}

//按社区查找
func GetCommunityPostIDsInOrder(p *models.ParamPostList) ([]string, error) {

	orderkey := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		orderkey = getRedisKey(KeyPostScoreZSet)
	}
	//使用zinterstore 把分区的帖子zset与分数帖子的zset生成一个新的zset
	//针对新的zset 按之前的逻辑去数据

	//利用缓存的key减少zinterstore执行次数
	key := orderkey + strconv.Itoa(int(p.CommunityID))
	if rdb.Exists(key).Val() < 1 {
		//不存在需要计算
		pipeline := rdb.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, getRedisKey(KeyCommunitySetPF+strconv.Itoa(int(p.CommunityID))), orderkey)
		pipeline.Expire(key, 60*time.Second)
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	return getIDsFormKey(key, p.Page, p.Size)

}

//GetPostVoteData 根据ids 查询每一个帖子的投票数据
func GetPostVoteData(ids []string) (data []int64, err error) {

	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPrefix + id)
	//	val := rdb.ZCount(key, "1", "1").Val()
	//	data = append(data, val)
	//}
	data = make([]int64, 0, len(ids))

	//使用pipeline 减少网络连接的次数
	pipeline := rdb.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return nil, err
}
