package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func banner() {
	print("===========\n")
	print("Duvi Server\n")
	print("===========\n")
}

var cliCount = 0

func handle(c net.Conn) {

	println("[INFO] New connection")

	for {

		dataFromCli0, err := bufio.NewReader(c).ReadString('\n')

		if err != nil {
			println("[ERR] Get client data")
			fmt.Println(err)
			return
		}

		dataFromCli := strings.TrimSpace(string(dataFromCli0))
		fmt.Println("[DATA]", dataFromCli)

		// SWITCH STRUCTURE
		//
		// BUG: c.Close() not accessable
		//
		// switch dataFromCli {

		// case "exit":

		// 	println("[INFO] Client disconnected")
		// 	cliCount -= 1

		// case "duvi":
		// 	c.Write([]byte(string("basil\n")))

		// default:
		// 	c.Write([]byte(string("\n")))

		// }

		// IF STRUCTURE
		//
		if dataFromCli == "exit" {

			println("[INFO] Client disconnected")
			cliCount -= 1
			break

		} else if dataFromCli == "duvi" {

			c.Write([]byte(string("basil\n")))

		} else if dataFromCli == "basil" {

			c.Write([]byte(string("duvi\n")))

		} else {
			// Default
			//
			// Sends "nothing"
			c.Write([]byte(string("\n")))
		}

		// Client Counter
		//
		// cliCounter := strconv.Itoa(count) + "\n"
		// c.Write([]byte(string(cliCounter)))

		// msgToCli := "DUVI" + "\n"
		// c.Write([]byte(string(msgToCli)))

		//

		// Send arbitrary data to client
		//
		// msgToCli := ""

		// print("DuP# ")
		// fmt.Scanln(&msgToCli)

		// msgToCli += "\n"
		// c.Write([]byte(string(msgToCli)))

		//

		// Default
		//
		// Sends nothing
		//c.Write([]byte(string("\n")))

	}

	// Works when if structure running
	c.Close()
}

func serve() {

	PORT := ""

	print("PORT: ")
	fmt.Scanln(&PORT)

	l, err := net.Listen("tcp4", fmt.Sprint(":", PORT))

	if err != nil {
		println("[ERR] Start listening")
		fmt.Println(err)
		return
	}

	defer l.Close()

	println("[INFO] Listening")

	for {

		c, err := l.Accept()

		if err != nil {
			println("[ERR] Accept client")
			fmt.Println(err)
			return
		}

		go handle(c)
		cliCount++

	}
}

func main() {
	banner()
	serve()
}
