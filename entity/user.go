package entity

import (
	"encoding/json"
	"github.com/cakemarketing/snowbank/stores/redis"
)

type User struct {
	name string
}

func (u *User) Name() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u User) Parse() ([]byte, error) {
	userJson, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	return userJson, nil
}
func UserExists(user User, db *redis.Redis) bool {
	_, err := db.Client.Get(user.name).Result()
	if err != nil {
		return false
	}
	return true
}
func GetUserEquations(user User, db *redis.Redis) (Equation, error) {
	equationsJson, _ := db.Client.Get(user.name).Result()
	var equations Equation
	e, err := ParseFromJson(equationsJson, equations)
	if err != nil {
		return Equation{}, err
	}
	return e, nil
}
func SaveUser(user User, equation Equation, db *redis.Redis) error {
	equationJson, err := equation.ParseToJson()
	if err != nil {
		return err
	}
	db.Client.Set(user.name, equationJson, 0)
	return nil
}
