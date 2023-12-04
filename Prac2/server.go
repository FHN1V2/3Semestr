	package main

	import (
		"fmt"
		"net"
		"strings"
		"sync"
	)

	type Database struct {
		Stack *Stack
		Queue *Queue
		set   *Set
		hmap  *HashMap
	}

	func NewDatabase() *Database {
		set := NewSet()
		return &Database{
			Stack: &Stack{},
			Queue: &Queue{},
			set:   set,
			hmap:  &HashMap{},
		}
	}

	func handleConnection(conn net.Conn, db *Database) {
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

			if len(parts) == 1 {
				command := parts[0]
				switch command {
				case "SPOP":
					poppedValue := db.Stack.Spop()
					conn.Write([]byte("Popped value from stack: " + poppedValue + "\n"))
				case "QPOP":
					poppedValue := db.Queue.Qpop()
					conn.Write([]byte("Popped value from queue: " + poppedValue + "\n"))
				case "SETPRINT":
					setContents := db.set.SetPrint()
					conn.Write([]byte(setContents + "\n"))
				}
			} else if len(parts) >= 2 {
				command := parts[0]
				switch command {

				case "SPUSH":

					value := parts[1]
					db.Stack.Spush(value)
					conn.Write([]byte("Value pushed to stack: " + value + "\n"))

				case "QPUSH":
					value := parts[1]
					db.Queue.Qpush(value)
					conn.Write([]byte("Value pushed to queue: " + value + "\n"))

				case "SETPUSH":
					value := parts[1]
					db.set.SetAdd(value)
					conn.Write([]byte("Value pushed to set: " + value + "\n"))

				case "SETDEL":
					value := parts[1]
					db.set.SetRemove(value)
				case "HPUSH":
					if len(parts) < 3 {
						conn.Write([]byte("Command must have at least 3 arguments\n"))
						continue
					}
					key, value := parts[1], parts[2]
					if db.hmap.Hadd(key, value) != nil {
						conn.Write([]byte("Key already exists in hashmap.\n"))
						continue
					}
					db.hmap.Hadd(key, value)

					conn.Write([]byte("Value pushed to hashtable: " + value + "\n"))
				case "HDEL":
					key := parts[1]
					db.hmap.Hdel(key)
				case "HGET":
					key := parts[1]
					poppedValue, err := db.hmap.Hget(key)
					if err != nil {
						conn.Write([]byte("Key not found in hash table.\n"))
						continue
					}
					conn.Write([]byte(key + " " + poppedValue + ".\n"))
				default:
					conn.Write([]byte("Unknown command: " + command + "\n"))
				}

			}
			if len(parts) >= 1 {
				//command := parts[0]
				//response := "Command received: " + command
				//conn.Write([]byte(response + "\n"))
			} else {
				conn.Write([]byte("Zero command \n"))
			}
		}
	}

	func main() {
		listen, err := net.Listen("tcp", "0.0.0.0:6379")
		if err != nil {
			fmt.Printf("Failed to bind to port: %s\n", err)
			return
		}

		defer listen.Close()

		fmt.Println("Server is listening on port 6379")

		var wg sync.WaitGroup
		db := NewDatabase()

		for {
			conn, err := listen.Accept()
			if err != nil {
				fmt.Printf("Failed to accept connection: %s\n", err)
				continue
			}

			wg.Add(1)
			go func() {
				defer wg.Done()
				handleConnection(conn, db)
			}()
			//wg.Wait()
		}


	}
