package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to TCP server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to connect to server:", err)
		return
	}
	defer conn.Close()

	// Read user input from stdin
	reader := bufio.NewReader(os.Stdin)

	for {
		// Read user input
		fmt.Print("Enter a message: ")
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read user input:", err)
			return
		}

		// Send message to server
		// _, err = fmt.Fprintf(conn, message)

		// if err != nil {
		// 	fmt.Println("Failed to send message:", err)
		// 	return
		// }
		println(fmt.Fprintf(conn, message))
		// Read server response
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Failed to read server response:", err)
			return
		}

		// Print server response
		fmt.Println("Server response:", response)
	}
}
