package db

import (
	"github.com/cakemarketing/snowbank/stores"
	"github.com/cakemarketing/snowbank/stores/redis"
	_ "github.com/go-redis/redis"
	redis2 "github.com/go-redis/redis"

	"sync"
)

type RedisCache struct {
	redis.Redis
	MU *sync.RWMutex
}

type RedisCacher interface {
	//	SMembers(key string)  *redis2.StringSliceCmd
	Scan(cursor uint64, match string, count int64) *redis2.ScanCmd
	SPop(key string) *redis2.StringCmd
	SAdd(key string, members ...interface{}) *redis2.IntCmd
}
type EquationAdder interface {
	SAdd(key string, members ...interface{}) *redis2.IntCmd
}

func (c RedisCache) SPop(key string) *redis2.StringCmd {
	return c.Client.SPop(key)

}
func (c RedisCache) SMembers(key string) *redis2.StringSliceCmd {
	cmd := c.Client.SMembers(key)
	return cmd
}
func (c RedisCache) Scan(cursor uint64, match string, count int64) *redis2.ScanCmd {
	cmd := c.Client.Scan(cursor, match, count)
	return cmd
}
func (c *RedisCache) SAdd(key string, members ...interface{}) *redis2.IntCmd {
	return c.SAdd(key, members...)
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
