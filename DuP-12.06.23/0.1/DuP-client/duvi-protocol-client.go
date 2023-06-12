package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func banner() {

	print("==========\n")
	print("DuP Client\n")
	print("==========\n")
}

func connect() {

	CONNECT := ""

	print("HOST:PORT ")
	fmt.Scanln(&CONNECT)

	c, err := net.Dial("tcp", CONNECT)

	if err != nil {
		println("[ERR] Could not connect to ", CONNECT)
		return
	}

	for {

		msgReader := bufio.NewReader(os.Stdin)

		// Data to send
		fmt.Print("DuP$ ")
		dataToSrv, _ := msgReader.ReadString('\n')

		// Send
		fmt.Fprintf(c, dataToSrv+"\n")

		// Get data
		dataFromSrv, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("DuP# " + dataFromSrv)

		// IF STRUCTURE
		//
		// if strings.TrimSpace(string(dataToSrv)) == "exit" {
		// 	println("[INFO] Exiting")
		// 	return
		// }

		switch strings.TrimSpace(string(dataToSrv)) {

		case "exit":

			println("\n[INFO] Exiting")
			os.Exit(0)
			return

		default:
			continue
		}
	}
}

func main() {
	banner()
	connect()
}
