package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

type Equation struct {
	Equations []EquationObject `json:"equations"`
}

type EquationObject struct {
	Expresion string    `json:"expresion"`
	Result    string    `json:"result"`
	Timestamp time.Time `json:"timestamp"`
}

func (e Equation) ParseToJson() ([]byte, error) {
	equationsJson, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return equationsJson, nil
}
func ParseFromJson(equationsJson string, e Equation) (Equation, error) {
	err := json.Unmarshal([]byte(equationsJson), &e)
	if err != nil {
		return Equation{}, fmt.Errorf("error while parsing")
	}
	return e, nil
}
