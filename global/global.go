package global

import (
	"ginEssential/config"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_REDIS  *redis.Client
)

