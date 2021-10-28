package calculator

import (
	"errors"
	"fmt"
	"github.com/apex/log"
	"os"
	"strconv"
)

func SolvePostfix(postfixString string) (float64, error) {
	log.Info(" SolvePostfix called")
	var stack FloatStack
	a := ""
	b := ""
	var aNum float64
	var bNum int = -1
	var aEmpty bool = true
	for index := 0; index < len(postfixString); index++ {
		element := postfixString[index]

		if IsOperand(postfixString[index]) {
			if aEmpty {
				a = a + string(element)
			} else {
				b = b + string(element)
			}
			i := 0
			for i = index + 1; postfixString[i] != ' '; i++ {
				if IsOperand(postfixString[i]) {
					if aEmpty {
						a = a + string(postfixString[i])
					} else {
						b = b + string(postfixString[i])
					}
				}
			}
			if postfixString[index+1] != ' ' {
				index = i
			}
			var err error
			if aEmpty {
				x, err := strconv.Atoi(a)
				if err != nil {
					// handle error
					fmt.Println(err)
					os.Exit(2)
				}
				aNum = float64(x)
				stack.Push(aNum)
			}
			aEmpty = false

			if len(b) > 0 {
				bNum, err = strconv.Atoi(b)
				stack.Push(float64(bNum))
				b = ""
			}
			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}

		} else if IsOperator(postfixString[index]) && postfixString[index] != ' ' {
			num2, ok := stack.Pop()
			if !ok {
				log.Error("invaild input Expression")
				return float64(0), errors.New("invaild input Expression")
			}

			num1, ok := stack.Pop()
			if !ok {
				log.Error("invaild input Expression")
				return float64(0), errors.New("invaild input Expression")
			}

			var err error
			operator := string(postfixString[index])
			aNum, err = SolveAB(float64(num1), float64(num2), operator)
			if err != nil {
				return float64(0), err
			}
			stack.Push(aNum)
			b = ""
			bNum = 0
		}
	}
	log.Info("SolvePostfix ended")
	return aNum, nil
}
func SolveAB(a float64, b float64, operator string) (float64, error) {
	if operator == "+" {
		return float64(a + b), nil
	} else if operator == "-" {
		return float64(a - b), nil
	} else if operator == "*" {
		return a * b, nil
	} else if operator == "/" {
		if b == 0 {
			log.Error("infinity any number can't be divided by zero (b=0)")
			return 0, errors.New("infinity any number can't be divided by zero")
		}
		c := a / b
		return float64(c), nil
	}
	return 0, nil
}
