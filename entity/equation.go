package entity

import (
	"challenge1/db"
	"encoding/json"
	"github.com/apex/log"
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
	SaveEquation(user User, redisdb db.EquationAdder) error
}

func (equation *Equation) SaveEquation(user User, redisdb db.EquationAdder) error {

	equationJson, err := json.Marshal(equation)
	if err != nil {
		return err
	}
	_, err = redisdb.SAdd(user.name, equationJson).Result()
	if err != nil {
		log.Error("erorr while sadd details =" + err.Error())
		return err
	}
	return nil
}
