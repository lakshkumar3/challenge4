package service

import (
	"challenge1/entity"
	"github.com/apex/log"
	"github.com/cakemarketing/snowbank/stores"
	"github.com/cakemarketing/snowbank/stores/redis"
	"github.com/cakemarketing/snowbank/warehouse"
)

func User(user entity.User, equation entity.EquationCollection) error {
	opts := &warehouse.Options{
		ClientId:          0,
		PreferredDatabase: []stores.DatabaseType{"redis"},
		Region:            "",
		TimeoutSeconds:    5,
	}
	redisinterface := warehouse.GetConnection("redis", opts)
	db, ok := redisinterface.(*redis.Redis)
	if !ok {
		log.Fatal("redis parse error")
	}

	entity.SaveUser(user, equation, db)

	return nil
}
