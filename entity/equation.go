package entity

import (
	"time"
)

type Equation struct {
	Expresion string    `json:"expresion"`
	Result    string    `json:"result"`
	Timestamp time.Time `json:"timestamp"`
}
