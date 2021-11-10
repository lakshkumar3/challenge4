package server

import (
	"bufio"
	"challenge1/calculator"
	"challenge1/entity"
	"challenge1/service"
	"fmt"
	"github.com/apex/log"
	"github.com/cakemarketing/go-common/v5/settings"
	"net"
	"os"
	"os/signal"
	"time"
)

func StartServer() error {
	var totalCount int = 0
	log.Info("program starting ")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	src := fmt.Sprintf(":" + settings.GetString("LOCAL_PORT"))
	listener, err := net.Listen("tcp", src)
	defer listener.Close()
	if err != nil {
		log.Fatal("could not listen to server " + err.Error())
		return err
	}
	fmt.Printf("Listening on %s.\n", src)

	go func() {
		<-ch
		log.Info("\n number of expression calculated by all clients" + fmt.Sprint(totalCount))
		log.Info("server closing it calculated total " + fmt.Sprint(totalCount) + " Expressions across all clients")
		listener.Close()
		os.Exit(1)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}
		serviceInfoMsg := " \nthis service calculates the expression  it can only perform 4 operations  + - / *"
		defer conn.Close()
		conn.Write([]byte("enter your name \n"))
		var name string
		fmt.Fscanln(conn, &name)
		conn.Write([]byte("Welcome " + name + serviceInfoMsg + "\n"))
		go handleConnection(conn, &totalCount, name)
	}
	return nil
}

func handleConnection(conn net.Conn, total_count *int, name string) {
	var client_count int = 0
	remoteAddr := conn.RemoteAddr().String()
	log.Info("Client connected from " + remoteAddr)
	scanner := bufio.NewScanner(conn)
	var equation entity.EquationCollection
	conn.Write([]byte("\nEnter Expression or type /quit to Exit : \n"))
	for {
		ok := scanner.Scan()

		if !ok {
			break
		}
		handleMessage(scanner.Text(), conn, &client_count, &equation)
	}
	*total_count = *total_count + client_count
	log.Info(remoteAddr + " " + name + " client calculate total " + fmt.Sprint(client_count) + " Expressions")
	log.Info("Client  " + name + " disconnected.")
	var user entity.User
	user.SetName(name)
	err := service.User(user, equation)
	if err != nil {
		log.Error("error at service" + err.Error())
		return
	}
	log.Info(fmt.Sprintf("user :%s  all equations %v", user.Name(), equation))
}

func handleMessage(message string, conn net.Conn, client_count *int, equation *entity.EquationCollection) {
	log.Info("> " + message)

	if len(message) > 0 {
		switch {
		case message == "/quit":
			conn.Write([]byte(" Total Expression count " + fmt.Sprint(*client_count) + "by client" + "\n"))
			log.Info("Quitting.")
			conn.Write([]byte("I'm shutting down now.\n"))
			log.Info("< " + "%quit%")
			conn.Write([]byte("%quit%\n"))
			conn.Close()

		default:
			result, err := calculator.CalculateExpression(message)
			if err != nil {
				log.Error(err.Error())
				conn.Write([]byte(err.Error() + "\n"))
			} else {
				*client_count++
				clientaddress := conn.RemoteAddr().String()
				equation.Equations = append(equation.Equations, entity.Equation{
					Expresion: message,
					Result:    result,
					Timestamp: time.Now(),
				})
				log.Info("client " + clientaddress + " " + result)
				conn.Write([]byte(result + "  Expression count " + fmt.Sprint(*client_count) + "\n"))
				conn.Write([]byte("Enter Expression or type /quit to Exit : \n"))
			}
		}
	}
}
