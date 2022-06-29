package redis

//redis key

const (
	KeyPrefix              = "bluebell"
	KeyPostTimeZSet        = ":post:time"   //zset 帖子发帖时间以及帖子
	KeyPostScoreZSet       = ":post:score"  //zset 帖子以及分数
	KeyPostVotedZSetPrefix = ":post:voted:" //zset 记录用户以及投票类型 参数是帖子id post_id
	KeyCommunitySetPF      = "community:"   //set保存每个分区下帖子的id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
