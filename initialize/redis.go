package initialize

import (
	"ginEssential/global"
	"github.com/go-redis/redis"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
	} else {
		global.GVA_REDIS = client
	}
}
