package service

import (
	"challenge1/entity"
	"github.com/apex/log"
	"github.com/cakemarketing/snowbank/stores/redis"
	"github.com/cakemarketing/snowbank/warehouse"
)

func User(user entity.User, equation entity.Equation) error {
	opts := &warehouse.Options{}
	redisinterface := warehouse.GetConnection("redis", opts)
	db, ok := redisinterface.(*redis.Redis)
	if !ok {
		log.Fatal("redis parse error")
	}

	entity.SaveUser(user, equation, db)

	return nil
}
