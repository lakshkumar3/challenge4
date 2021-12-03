package service

import (
	"challenge1/db"
	"challenge1/entity"
	"github.com/apex/log"
	"github.com/cakemarketing/snowbank/warehouse"
	_ "github.com/go-sql-driver/mysql"
)

var (
	userService SaveService
)

func init() {
	userService = &UserService{}
}

type UserService struct {
}
type SaveService interface {
	SaveCache() error
	SaveEquation(user entity.User, equation entity.Equation) error
}

func (*UserService) SaveEquation(user entity.User, equation entity.Equation) error {
	opts := &warehouse.Options{}
	redisinterface := warehouse.GetConnection("redis", opts)
	redisdb, ok := redisinterface.(*db.RedisCache)
	if !ok {
		log.Fatal("redis parse error")
	}
	err := equation.SaveEquation(user, redisdb)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("sucessfullly updated on redis")
	return nil
}
func (*UserService) SaveCache() error {
	opts := &warehouse.Options{}
	msSqlinterface := warehouse.GetConnection("aurora", opts)
	sqldb, ok := msSqlinterface.(*db.Mysql)
	if !ok {
		log.Error("error while parsing")
	}
	err := sqldb.Ping()
	if err != nil {
		log.Error(err.Error())
	}
	redisinterface := warehouse.GetConnection("redis", opts)
	redisdb, ok := redisinterface.(*db.RedisCache)
	if !ok {
		log.Fatal("redis parse error")
	}
	var cache entity.Cache
	err = cache.SaveEquationCache(redisdb, sqldb)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("sucessfullly updated on redis")
	return nil
}
