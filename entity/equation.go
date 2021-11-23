package entity

import (
	"challenge1/db"
	"encoding/json"
	"github.com/apex/log"
	"github.com/cakemarketing/snowbank/warehouse"
	"time"
)

var EquationEntity Saver

type Equation struct {
	Expresion string    `json:"expresion"`
	Result    string    `json:"result"`
	Timestamp time.Time `json:"timestamp"`
}

func init() {
	EquationEntity = &Equation{}
}

type Saver interface {
	SaveEquation(user User) error
}

func (equation *Equation) SaveEquation(user User) error {
	opts := &warehouse.Options{}
	redisinterface := warehouse.GetConnection("redis", opts)
	redisdb, ok := redisinterface.(*db.RedisCache)
	if !ok {
		log.Fatal("redis parse error")
	}
	equationJson, err := json.Marshal(equation)
	if err != nil {
		return err
	}
	redisdb.MU.Lock()
	cmd := redisdb.Client.SAdd(user.name, equationJson)
	if cmd.Err() != nil {
		log.Error("erorr while sadd details =" + cmd.Err().Error())
		redisdb.MU.Unlock()
		return err
	}
	redisdb.MU.Unlock()
	return nil
}
