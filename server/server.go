package server

import (
	"bufio"
	"challenge1/calculator"
	"fmt"
	"github.com/apex/log"
	"github.com/cakemarketing/go-common/v5/settings"
	"net"
	"os"
	"os/signal"
	"strconv"
)

func server() error {
	var totalCount int = 0
	log.Info("program exiting ")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	src := ":" + strconv.Itoa(settings.GetInt("LOCAL_PORT"))
	listener, err := net.Listen("tcp", src)
	if err != nil {
		log.Fatal("could not listen to server " + err.Error())
		return err
	}
	fmt.Printf("Listening on %s.\n", src)

	go func() {
		<-ch
		fmt.Println("\n number of expression calculated by all clients", totalCount)
		log.Info("server closing it calculated total " + strconv.Itoa(totalCount) + " Expressions across all clients")
		listener.Close()
		os.Exit(1)
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}
		serviceInfoMsg := "Welcome this service calculates the expression  it can only perform 4 operations  + - / *"
		conn.Write([]byte(serviceInfoMsg))
		go handleConnection(conn, &totalCount)
		defer conn.Close()
	}
	return nil
}

func handleConnection(conn net.Conn, total_count *int) {
	var client_count int = 0
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from " + remoteAddr)
	log.Info("Client connected from " + remoteAddr)
	scanner := bufio.NewScanner(conn)

	for {
		conn.Write([]byte("\nEnter Expresstion or type /quit to Exit : "))
		ok := scanner.Scan()

		if !ok {
			break
		}

		handleMessage(scanner.Text(), conn, &client_count)

	}
	*total_count = *total_count + client_count
	log.Info(remoteAddr + " client calculate total " + strconv.Itoa(client_count) + "Expressions")
	fmt.Println("Client at " + remoteAddr + " disconnected.")
}

func handleMessage(message string, conn net.Conn, client_count *int) {
	fmt.Println("> " + message)

	if len(message) > 0 {
		switch {
		case message == "/quit":
			conn.Write([]byte(" Total Expression count " + strconv.Itoa(*client_count) + "by client" + "\n"))
			fmt.Println("Quitting.")
			conn.Write([]byte("I'm shutting down now.\n"))
			fmt.Println("< " + "%quit%")
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

				log.Info("client " + clientaddress + " " + result)
				conn.Write([]byte(result + "  Expression count " + strconv.Itoa(*client_count) + "\n"))
			}
		}
	}
}
func Start() error {
	err := server()
	return err

}
