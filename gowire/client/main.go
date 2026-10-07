package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:6140")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	fmt.Println("Connecting to localhost:6140...")

	buffer := bufio.NewReader(os.Stdin)
	serverResponse := bufio.NewReader(conn)
	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println(err)
			return
		}

		response, err := serverResponse.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("~", response)
		time.Sleep(time.Millisecond * 500)
		if response == "BYE\n" {
			return
		}
	}
}
