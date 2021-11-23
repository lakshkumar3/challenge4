package service

import (
	"challenge1/entity"
	"github.com/apex/log"
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

	err := equation.SaveEquation(user)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("sucessfullly updated on redis")
	return nil
}
func (*UserService) SaveCache() error {
	var cache entity.Cache
	err := cache.SaveEquationCache()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	log.Info("sucessfullly updated on redis")
	return nil
}
