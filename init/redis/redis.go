package redis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/lvjiaben/go-wheel/init/viper"
)

var rdb *redis.Client

func Load() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.Conf.Redis.Host, viper.Conf.Redis.Port),
		Password: viper.Conf.Redis.Pass,
		DB:       viper.Conf.Redis.Db,
		PoolSize: viper.Conf.Redis.PoolSize,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func Close() {
	_ = rdb.Close()
}
