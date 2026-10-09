package main

import (
	"bufio"
	"fmt"
	"net"
	"server/keyValues"
	"strings"
)

func main() {
	kv := keyvalues.CreateKeyValue()
	manager := NewConnectionManager()
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

		client := manager.NewClient(conn)
		manager.Add(client)
		fmt.Printf(
			"Client %d connected: %s\n",
			client.ID,
			conn.RemoteAddr(),
		)
		fmt.Println(client.IsAuthenticated())
		fmt.Printf("Active clients: %d\n", manager.Count())

		go func() {
			handleConnection(client, manager, kv)
		}()
	}
}

func handleConnection(client *Client, manager *ConnectionManager,kv *keyvalues.KeyValues) {
	defer func() {
		manager.Remove(client.ID)
		client.Conn.Close()

		fmt.Printf("Client %d disconnected\n", client.ID)
		fmt.Printf("Active clients: %d\n", manager.Count())
	}()

	conn := client.Conn

	buffer := bufio.NewReader(conn)

	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Printf(
				"Client %d read error: %v\n",
				client.ID,
				err,
			)
			return
		}
		parts := strings.Fields(line)
		if len(parts) == 0 {
			_, err = conn.Write([]byte("-ERR : Empty command" + "\n"))
			if err != nil {
				fmt.Println("Write error:", err)
				return
			}
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

		if !client.IsAuthenticated() {
			if strings.ToUpper(parts[0]) != "AUTH" {
			_, err := conn.Write(
				[]byte("-ERR authentication required\n"),
			)
			if err != nil {
				fmt.Println("Write error:", err)
				return
			}
			continue
		}
	}
		if strings.ToUpper(parts[0]) != "AUTH"{
			if len(parts) != 3 {
				_, err := conn.Write(
				[]byte("-ERR AUTH require 3 parameters\n"),
			)
			if err != nil {
				fmt.Println("Write error:", err)
				return
			}
			}
			
		}

		if err := validate_string(parts); err != nil {
			_, writeErr := conn.Write(
				[]byte("-ERR " + err.Error() + "\n"),
			)

			if writeErr != nil {
				fmt.Println("Write error:", writeErr)
				return
			}

			continue
		}

		response := dispatch(parts, kv)
		_, err = conn.Write([]byte(response))

		if err != nil {
			fmt.Printf("Client %d write error: %v\n", client.ID, err)
			return
		}
	}
}
