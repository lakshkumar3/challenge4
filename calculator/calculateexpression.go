package calculator

import (
	"fmt"
	"github.com/apex/log"
)

func CalculateExpression(inputExpression string) (float64, error) {
	log.Debug("CalculateExpression started ")

	postfixString := ToPostfix(inputExpression)
	result, err := SolvePostfix(postfixString)
	if err != nil {
		log.Error("CalculateExpression " + err.Error())
		return 0, err
	} else {
		fmt.Printf(inputExpression + " = ")
		return result, nil
	}
}
