package entity

import (
	"encoding/json"
	"github.com/apex/log"
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

func SaveUser(user User, equation EquationCollection, db *redis.Redis) error {
	for _, element := range equation.Equations {
		equationJson, err := json.Marshal(element)
		if err != nil {
			return err
		}
		cmd := db.Client.SAdd(user.name, equationJson)
		if cmd.Err() != nil {
			log.Error("erorr while sadd details =" + cmd.Err().Error())
		}
	}
	return nil
}
