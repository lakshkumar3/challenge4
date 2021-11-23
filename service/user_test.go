package service

import (
	"challenge1/entity"
	"fmt"
	"testing"
)

var (
	SaveCacheFunc    func() error
	SaveEquationFunc func(user entity.User, equation entity.Equation) error
)

type ServiceMock struct{}

func (m *ServiceMock) SaveCache() error {
	return SaveCacheFunc()
}
func (m *ServiceMock) SaveEquation(user entity.User, equation entity.Equation) error {

	return SaveEquationFunc(user, equation)
}

func TestUserService_SaveCache(t *testing.T) {
	userService = &ServiceMock{}
	SaveCacheFunc = func() error {
		return nil
	}
	err := userService.SaveCache()
	if err != nil {
		t.Fail()
	}

}

func TestNegativeUserService_SaveCache(t *testing.T) {
	userService = &ServiceMock{}
	SaveCacheFunc = func() error {
		return fmt.Errorf("error 404")
	}
	err := userService.SaveCache()
	if err == nil {
		t.Fail()
	}

}
func TestUserService_SaveEquation(t *testing.T) {
	userService = &ServiceMock{}
	SaveEquationFunc = func(user entity.User, equation entity.Equation) error {
		return nil
	}
	user := entity.User{}
	equation := entity.Equation{}
	err := userService.SaveEquation(user, equation)
	if err != nil {
		t.Fail()
	}

}

func TestNegativeUserService_SaveEquation(t *testing.T) {
	userService = &ServiceMock{}
	SaveEquationFunc = func(user entity.User, equation entity.Equation) error {
		return fmt.Errorf("error 404")
	}
	user := entity.User{}
	equation := entity.Equation{}
	err := userService.SaveEquation(user, equation)
	if err != nil {
		t.Fail()
	}

}
