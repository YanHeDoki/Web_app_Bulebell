package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 //每一票占432分
)

var (
	ErrVoteTimeExpire = errors.New("超出投票时间")
	ErrVoteRepset     = errors.New("已投票过")
)

func VoteForPost(Userid, postId string, direction float64) error {

	//1.判断投票的限制

	//从redis获取帖子发布的时间
	posttime := rdb.ZScore(getRedisKey(KeyPostTimeZSet), postId).Val()

	if float64(time.Now().Unix())-posttime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}

	//更新分数
	//先查当前用户对当前帖子的投票记录
	uservote := rdb.ZScore(getRedisKey(KeyPostScoreZSet+postId), Userid).Val()

	// 更新 如果和上次投票保持一致 应到不允许投票
	if direction == uservote {
		return ErrVoteRepset
	}
	var op float64
	if direction > uservote {
		op = 1
	} else {
		op = -1
	}

	diff := math.Abs(uservote - direction) //计算两次投票的差值
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postId)

	//记录用户为该帖子投票
	if direction == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSetPrefix+postId), Userid)

	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSetPrefix+postId), redis.Z{
			Score:  direction, //当前用户投的什么票
			Member: Userid,
		})
	}
	_, err := pipeline.Exec()

	return err
}

func CreatePost(postId, CommunityID int64) error {

	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{Score: float64(time.Now().Unix()), Member: postId})
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{Score: float64(time.Now().Unix()), Member: postId})
	ckey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(CommunityID)))
	pipeline.SAdd(ckey, postId)
	_, err := pipeline.Exec()

	return err
}
