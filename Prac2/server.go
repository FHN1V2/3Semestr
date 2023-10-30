package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	startServer()
}

type Database struct {
	stack  Stack
	queue  Queue
	hash   HashMap
	set    *Set // Change the field type to *Set
}


func startServer() {
	listen, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		return
	}
	defer listen.Close()

	fmt.Println("Server is listening on port 6379")

	db := Database{
		stack:  Stack{},
		queue:  Queue{},
		hash:   HashMap{},
		set:    NewSet(),
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		go handleConnection(conn, &db)
	}
}

func handleConnection(conn net.Conn, db *Database) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Connection error: %v\n", err)
			return
		}

		command := string(buf[:n])
		command = strings.TrimSpace(command)

		args := strings.Fields(command)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "SPUSH":
			if len(args) == 2 {
				db.stack.Spush(args[1])
				conn.Write([]byte("OK\n"))
			} else {
				conn.Write([]byte("Error: Invalid arguments\n"))
			}
		case "SPOP":
			value, err := db.stack.Spop()
			if err != nil {
				conn.Write([]byte("Error: " + err.Error() + "\n"))
			} else {
				conn.Write([]byte(value + "\n"))
			}
			
		case "QPUSH":
			if len(args) == 2 {
				db.queue.Qadd(args[1])
				conn.Write([]byte("OK\n"))
			} else {
				conn.Write([]byte("Error: Invalid arguments\n"))
			}
		case "QPOP":
			value := db.queue.Qdell()
			conn.Write([]byte(value + "\n"))
		case "HADD":
			if len(args) == 3 {
				db.hash.Hadd(args[1], args[2])
				conn.Write([]byte("OK\n"))
			} else {
				conn.Write([]byte("Error: Invalid arguments\n"))
			}
		case "HGET":
			if len(args) == 2 {
				value := db.hash.Hget(args[1])
				conn.Write([]byte(value + "\n"))
			} else {
				conn.Write([]byte("Error: Invalid arguments\n"))
			}
		case "HDEL":
			if len(args) == 2 {
				db.hash.Hdel(args[1])
				conn.Write([]byte("OK\n"))
			} else {
				conn.Write([]byte("Error: Invalid arguments\n"))
			}
		case "SETPUSH":
			if len(args) == 2 {
				db.set.SetAdd(args[1])
				conn.Write([]byte("OK\n"))
			} else {
				conn.Write([]byte("Error: Invalid arguments\n"))
			}
		case "SETDEL":
			if len(args) == 2 {
				db.set.SetRemove(args[1])
				conn.Write([]byte("OK\n"))
			} else {
				conn.Write([]byte("Error: Invalid arguments\n"))
			}
		case "SETPRINT":
			values := db.set.SetPrint()
			conn.Write([]byte(strings.Join(values, ",") + "\n"))
		default:
			conn.Write([]byte("Error: Unknown command\n"))
		}
	}
}
