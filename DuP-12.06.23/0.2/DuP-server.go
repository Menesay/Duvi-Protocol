package main

import (
	"fmt"
	"net"
)

func banner() {
	print("===========\n")
	print("Duvi Server\n")
	print("===========\n")
}

var cliCount = 0

func handle(c net.Conn) {

	defer c.Close()

	println("[INFO] New connection")
	println("[INFO] Client count:", cliCount)

	for {

		buffer := make([]byte, 1024)
		bytesRead, err := c.Read(buffer)

		if err != nil {
			fmt.Println("[ERR] Get client data\n", err)
			return
		}

		dataFromCli := string(buffer[:bytesRead])
		fmt.Println("[DATA]", dataFromCli)

		//

		dataToCli := ""

		if dataFromCli == "exit" {

			println("[INFO] Client disconnected")
			cliCount -= 1
			println("[INFO] Client count:", cliCount)
			break

		} else if dataFromCli == "duvi" {

			dataToCli = "basil"
			c.Write([]byte(dataToCli))
			dataToCli = ""

		} else if dataFromCli == "basil" {

			dataToCli = "duvi"
			c.Write([]byte(dataToCli))
			dataToCli = ""

		} else {
			// Default
			//
			// It must send "any char". In this case, a space.
			c.Write([]byte(" "))
		}

		_, err = c.Write([]byte(dataToCli))

		if err != nil {
			fmt.Println("[ERR] Send data to client\n", err)
			return
		}

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
		// if dataFromCli == "exit" {

		// 	println("[INFO] Client disconnected")
		// 	cliCount -= 1
		// 	break

		// } else if dataFromCli == "duvi" {

		// 	c.Write([]byte(string("basil\n")))

		// } else if dataFromCli == "basil" {

		// 	c.Write([]byte(string("duvi\n")))

		// } else {
		// 	// Default
		// 	//
		// 	// Sends "nothing"
		// 	c.Write([]byte(string("\n")))
		// }

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

		// Works when if structure running
	}
}

func serve() {

	PORT := ""

	print("PORT: ")
	fmt.Scanln(&PORT)

	l, err := net.Listen("tcp", fmt.Sprint(":", PORT))
	if err != nil {
		fmt.Println("[ERR] Start listening\n", err)
		return
	}

	println("[INFO] Listening")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("[ERR] Accept client\n", err)
			continue
		}

		go handle(conn)
		cliCount++

	}
}

func main() {
	banner()
	serve()
}
