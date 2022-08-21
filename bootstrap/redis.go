package bootstrap

import (
	"fmt"
	"go-mall-api/pkg/config"
	"go-mall-api/pkg/redis"
)

//SetupRedis 初始化 Redis
func SetupRedis() {
	//建立Redis连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
