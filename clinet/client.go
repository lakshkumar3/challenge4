package main

import (
	"bufio"
	"fmt"
	"github.com/apex/log"
	"github.com/cakemarketing/go-common/v5/settings"
	"net"
	"os"
	"regexp"
	"strconv"
)

func Client() {

	dest := ":" + strconv.Itoa(settings.GetInt("LOCAL_PORT"))

	conn, err := net.Dial("tcp", dest)
	if err != nil {
		if _, t := err.(*net.OpError); t {
			log.Fatal("Some problem connecting.")
		} else {
			log.Fatal("Unknown error: " + err.Error())
		}
		os.Exit(1)
	}

	go readConnection(conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(text))
		if err != nil {
			log.Error("Error writing to stream.")
			break
		}
	}
}
func readConnection(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)

		for {
			ok := scanner.Scan()
			if !ok {
				break
			}
			text := scanner.Text()

			command := handleCommands(text)
			if !command {
				fmt.Println(text)
			}

			if !ok {
				fmt.Println("Reached EOF on server connection.")
				break
			}
		}
	}
}

func handleCommands(text string) bool {
	r, err := regexp.Compile("^%.*%$")
	if err == nil {
		if r.MatchString(text) {

			switch {
			case text == "%quit%":
				log.Info("\b\bServer is leaving. Hanging up.")
				os.Exit(0)
			}
			return true
		}
	}
	return false
}
