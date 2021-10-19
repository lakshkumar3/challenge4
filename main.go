package main

import (
	"fmt"
	"os"
	"os/signal"
	_ "github.com/apex/log"

)


func main() {
	ch := make(chan os.Signal, 1)
	fmt.Print("Enter  expression: ")
	/*	infixString, err := ReadFromInput()
		if err != nil {
			fmt.Println("Error when scanning input:", err.Error())
			return
		}*/
	var inputExpression string
	fmt.Scanln(&inputExpression)


	postfixString := ToPostfix(inputExpression)
	result := SolvePostfix(postfixString)
	fmt.Println("calculating")
	fmt.Println("Press ctrl+C to see result")

	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("calculating")
	fmt.Println(result)


}
