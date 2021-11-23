package entity

import (
	"fmt"
	"testing"
)

var (
	SaveEquationfunc func(user User) error
)

type EquationEntityMock struct{}

func (e EquationEntityMock) SaveEquation(user User) error {
	return SaveEquationfunc(user)
}

func TestEquation_SaveEquation(t *testing.T) {
	equationEntity := &EquationEntityMock{}
	SaveEquationfunc = func(user User) error {
		return nil
	}
	user := User{}
	user.SetName("laksh")
	err := equationEntity.SaveEquation(user)
	if err != nil {
		t.Fail()
	}

}
func TestNegativeEquation_SaveEquation(t *testing.T) {
	equationEntity := &EquationEntityMock{}
	SaveEquationfunc = func(user User) error {
		return fmt.Errorf("error 404")
	}
	user := User{}
	user.SetName("laksh")
	err := equationEntity.SaveEquation(user)
	if err == nil {
		t.Fail()
	}

}
