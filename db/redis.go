package db

import (
	"github.com/cakemarketing/snowbank/stores"
	"github.com/cakemarketing/snowbank/stores/redis"
	"sync"
)

type RedisCache struct {
	redis.Redis
	MU *sync.RWMutex
}

func ConnectRedis(opts *redis.Options) *RedisCache {
	redisdb := redis.ConnectRedis(opts)
	db := &RedisCache{
		Redis: redis.Redis{Client: redisdb.Client},
		MU:    &sync.RWMutex{},
	}
	return db
}
func (*RedisCache) GetModuleType() stores.ModuleType {
	return stores.ModuleType_DB
}

func (*RedisCache) GetDatabaseType() stores.DatabaseType {
	return stores.DatabaseType_Aurora
}

func (db *RedisCache) Healthy() error {
	return db.Healthy()
}
