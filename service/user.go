package service

import (
	"challenge1/entity"
	"github.com/apex/log"
	"github.com/cakemarketing/snowbank/stores"
	"github.com/cakemarketing/snowbank/stores/redis"
	"github.com/cakemarketing/snowbank/warehouse"
)

func User(user entity.User, equation entity.Equation) error {
	opts := &warehouse.Options{
		ClientId:          0,
		PreferredDatabase: []stores.DatabaseType{"redis"},
		Region:            "",
		TimeoutSeconds:    5,
	}
	redisinterface := warehouse.GetConnection("redis", opts)
	db, ok := redisinterface.(*redis.Redis)
	if !ok {
	}
	ok = entity.UserExists(user, db)
	if ok {
		oldEquations, err := entity.GetUserEquations(user, db)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		var allEquations entity.Equation
		allEquations.Equations = append(oldEquations.Equations, equation.Equations...)
		entity.SaveUser(user, allEquations, db)
	} else {
		entity.SaveUser(user, equation, db)
	}
	return nil
}
