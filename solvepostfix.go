package main

import (
	"fmt"
	"os"
	"strconv"
)

func SolvePostfix(postfixString string) float64 {
	a := ""
	b := ""
	var aNum float64
	var bNum int
	var aEmpty bool = true
	for index := 0; index < len(postfixString); index++ {
		element := postfixString[index]

		if IsOperand(postfixString[index]) {
			if aEmpty {
				a = a + string(element)
			} else {
				b = b + string(element)
			}
			for i := index + 1; postfixString[i] != ' '; i++ {
				if IsOperand(postfixString[i]) {
					if aEmpty {
						a = a + string(postfixString[i])
					} else {
						b = b + string(postfixString[i])
					}
				}
				index = i
				var err error
				if aEmpty {
					x, err := strconv.Atoi(a)
					if err != nil {
						// handle error
						fmt.Println(err)
						os.Exit(2)
					}
					aNum = float64(x)

				}
				aEmpty = false

				if len(b) > 0 {
					bNum, err = strconv.Atoi(b)
					b = ""
				}
				if err != nil {
					// handle error
					fmt.Println(err)
					os.Exit(2)
				}

			}

		} else if IsOperator(postfixString[index]) && postfixString[index] != ' ' {
			aNum = SolveAB(float64(aNum), float64(bNum), string(postfixString[index]))
			b = ""
			bNum = 0
		}
	}
	return aNum
}
func SolveAB(a float64, b float64, operator string) float64 {
	if operator == "+" {
		return float64(a + b)
	} else if operator == "-" {
		return float64(a - b)
	} else if operator == "*" {
		return a * b
	} else if operator == "/" {
		c := a / b
		return float64(c)
	}
	return 0
}
