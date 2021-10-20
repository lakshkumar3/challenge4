package main

import (
	"fmt"
	"os"
	"os/signal"
	"github.com/apex/log"
)


func main() {
	log.Info("main started ")

	ch := make(chan os.Signal, 1)
var count int =0
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		fmt.Println("\n number of expression calculated",count)
		log.Info("program exiting ")
		os.Exit(1)
	}()
	for i := 0; ; i++ {
	fmt.Print("Enter  expression: ")
	var inputExpression string
	fmt.Scanln(&inputExpression)
	postfixString := ToPostfix(inputExpression)
	result,err := SolvePostfix(postfixString)
		if err != nil {
			log.Error(err.Error())
			fmt.Println("error occured  ",err.Error())
		}else {
			count++
			fmt.Printf(inputExpression+" = ")
			fmt.Println(result)
		}
	fmt.Println("Press ctrl+C to exit")

}


}
