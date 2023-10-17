package models

type RedisInfo struct {
	Rediskey string
	Value    map[int64]interface{}
}
