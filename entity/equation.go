package entity

import (
	"encoding/json"
	"fmt"
	"time"
)

type EquationCollection struct {
	Equations []Equation `json:"equations"`
}

type Equation struct {
	Expresion string    `json:"expresion"`
	Result    string    `json:"result"`
	Timestamp time.Time `json:"timestamp"`
}

func (e EquationCollection) ParseToJson() ([]byte, error) {
	jsonEquations := make([][]byte, len(e.Equations))

	for _, equation := range e.Equations {
		jsonEquation, err := json.Marshal(equation)
		if err != nil {
			return nil, err
		}
		jsonEquations = append(jsonEquations, jsonEquation)
	}
	equationsJson, err := json.Marshal(e.Equations)
	if err != nil {
		return nil, err
	}
	return equationsJson, nil
}
func ParseFromJson(equationsJson string, e EquationCollection) (EquationCollection, error) {
	err := json.Unmarshal([]byte(equationsJson), &e)
	if err != nil {
		return EquationCollection{}, fmt.Errorf("error while parsing")
	}
	return e, nil
}
