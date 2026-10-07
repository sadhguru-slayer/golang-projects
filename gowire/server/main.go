package main

import (
	"bufio"
	"fmt"
	"net"
	keyvalues "server/keyValues"
	"strings"
)

func main() {
	kv := keyvalues.CreateKeyValue()
	listen, err := net.Listen("tcp", ":6140")

	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}

	defer listen.Close()
	fmt.Println("Listening to port :6140")

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Client disconnected. ERROR:", err)
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr(), conn.LocalAddr())
		go func() {
			handleConnection(conn, kv)
		}()
	}
}

func handleConnection(conn net.Conn, kv *keyvalues.KeyValues) {
	defer conn.Close()

	buffer := bufio.NewReader(conn)

	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected. ERROR:", err)
			return
		}
		parts := strings.Fields(line)
		if len(parts) == 0 {
			conn.Write([]byte("-ERR : Empty command" + "\n"))
			continue
		}
		switch strings.ToUpper(parts[0]) {
		case "PING":
			conn.Write([]byte("PONG\n"))
			continue
		case "QUIT":
			_, err := conn.Write([]byte("BYE\n"))
			if err != nil {
				fmt.Println("failed to send BYE:", err)
			}
			return
		}
		err = validate_string(parts)

		if err != nil {
			conn.Write([]byte("-ERR " + err.Error() + "\n"))
			continue
		}

		fmt.Println(line)
		response := dispatch(parts, kv)
		_, err = conn.Write([]byte(response))

		if err != nil {
			fmt.Println("CLient Disconnected")
			return
		}
	}
}
