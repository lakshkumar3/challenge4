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
	log.Info(fmt.Sprintf("Listening on %s.\n", src))

	go func() {
		<-ch
		log.Info("\n number of expression calculated by all clients" + fmt.Sprint(totalCount))
		log.Info(fmt.Sprintf("server closing it calculated total  %s Expressions across all clients", totalCount))
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
	var user entity.User
	user.SetName(name)
	var client_count int = 0
	remoteAddr := conn.RemoteAddr().String()
	log.Info("Client connected from " + remoteAddr)
	scanner := bufio.NewScanner(conn)
	conn.Write([]byte("\nEnter Expression or type /quit to Exit : \n"))
	for scanner.Scan() {

		equation, err := handleMessage(scanner.Text(), conn, &client_count)
		if err == nil {
			msg := fmt.Sprintf("client expression   client=%s equation=%s result=%s", name, equation.Expresion, equation.Result)
			log.Info(msg)
			err := service.User(user, equation)
			if err != nil {
				log.Error(err.Error())
			}
		} else if err.Error() == "quit" {
			conn.Close()
			break
		} else {
			log.Info("wrong message ")
		}
	}
	*total_count = *total_count + client_count
	log.Info("connection disconnected client " + name + " client calculated total " + fmt.Sprint(client_count) + " Expressions")
}

func handleMessage(message string, conn net.Conn, client_count *int) (entity.Equation, error) {
	log.Info("> " + message)
	if len(message) > 0 {
		switch {
		case message == "/quit":
			conn.Write([]byte(" Total Expression count " + fmt.Sprint(*client_count) + "by client" + "\n"))
			log.Info("Quitting.")
			conn.Write([]byte("I'm shutting down now.\n"))
			log.Info("< " + "%quit%")
			conn.Write([]byte("%quit%\n"))
			return entity.Equation{}, fmt.Errorf("quit")

		default:
			result, err := calculator.CalculateExpression(message)
			if err != nil {
				log.Error(err.Error())
				conn.Write([]byte(err.Error() + "\n"))
			} else {
				strResult := fmt.Sprint(result)
				*client_count++
				clientaddress := conn.RemoteAddr().String()

				log.Info("client " + clientaddress + " " + strResult)
				conn.Write([]byte(fmt.Sprintf("expression=%s result=%s  Expression count %d \n", message, strResult, *client_count)))
				conn.Write([]byte("Enter Expression or type /quit to Exit : \n"))
				return entity.Equation{
					Expresion: message,
					Result:    strResult,
					Timestamp: time.Now(),
				}, nil
			}
		}
	}
	return entity.Equation{}, fmt.Errorf("wrong message")
}
