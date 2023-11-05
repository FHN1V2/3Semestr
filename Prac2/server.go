package main

import (
	"fmt"
	"net"
	"sync"
	"strings"
)


func handleConnection(conn net.Conn, stack *Stack) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Connection closed: %s\n", err)
			return
		}

		input := string(buf[:n])
		fmt.Printf("Received input: %s\n", input)

		parts := strings.Fields(input)
		command := parts[0]

		if command == "SPUSH" {
			value := parts[1]
			stack.Spush(value)
			conn.Write([]byte("Value pushed to stack: " + value + "\n"))
		} else if command == "SPOP" {
			poppedValue := stack.Spop()
			conn.Write([]byte("Popped value from stack: " + poppedValue + "\n"))
		} else {
			conn.Write([]byte("Unknown command: " + command + "\n"))
		}

		response := "Command received: " + command
		conn.Write([]byte(response + "\n"))
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:6978")
	if err != nil {
		fmt.Printf("Failed to bind to port: %s\n", err)
		return
	}

	defer listen.Close()

	fmt.Println("Server is listening on port 6978")

	var wg sync.WaitGroup
	stack := &Stack{}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Failed to accept connection: %s\n", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			handleConnection(conn, stack)
		}()
	}
}
