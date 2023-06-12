package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

func banner() {
	print("===========\n")
	print("Duvi Server\n")
	print("===========\n")
}

//
//	Schema
//

func listDir(path string) (arr []string, err error) {

	filesArr := []string{}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		return filesArr, err
	}

	for _, file := range files {

		fileName := file.Name()

		if file.IsDir() {

			// fmt.Printf("D: %s\n", fileName)
			filesArr = append(filesArr, fileName)
		} else {
			// fmt.Printf("F: %s\n", fileName)
			filesArr = append(filesArr, fileName)
		}
	}
	return filesArr, nil

}

//
////////
//

var crrPth, _ = os.Getwd()
var cliCount = 0

func handle(c net.Conn) {

	cliIPPORT := c.RemoteAddr().String()

	defer c.Close()

	println("[INFO] New connection from:", cliIPPORT)
	println("[INFO] Client count:", cliCount)

	for {

		buffer := make([]byte, 1024)
		bytesRead, err := c.Read(buffer)

		if err != nil {
			fmt.Println("[ERR] Get client data\n", err)
			return
		}

		dataFromCli := string(buffer[:bytesRead])
		fmt.Println("[DATA]", cliIPPORT, ":", dataFromCli)

		//

		dataToCli := ""

		if dataFromCli == "exit" {

			println("[INFO]", cliIPPORT, "disconnected")
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

		} else if dataFromCli == "." {
			//
			// Send current path list
			//

			// filesFromArr, err := listDir(crrPth0)
			// if err != nil{

			// 	dataToCli = fmt.Sprintln(err)
			// 	c.Write([]byte(dataToCli))
			// }else{
			// 	for _, file := range filesFromArr{

			// 	}
			// }

			files, err := ioutil.ReadDir(crrPth)

			if err != nil {

				dataToCli = fmt.Sprintln(err)
				c.Write([]byte(dataToCli))

			} else {

				// Send C:\path:
				dataToCli = "\n" + crrPth + "#\n"

				fmt.Printf(crrPth + "#\n")
				for _, file := range files {

					fileName := file.Name()

					if file.IsDir() {

						fmt.Printf("D: %s\n", fileName)

						dataToCli += fmt.Sprintf("D: %s\n", fileName)

					} else {

						fmt.Printf("F: %s\n", fileName)

						dataToCli += fmt.Sprintf("F: %s\n", fileName)

					}
				}

				c.Write([]byte(dataToCli))
			}

		} else if dataFromCli == ".." {

			tmpDir := ""
			dataToCli = ""

			// C:\new\old --> C:\new
			// It changes directory to previous
			for i, letter := range crrPth {
				if string(letter) == "/" || string(letter) == "\\" {
					tmpDir = crrPth[:i]

				}

			}

			//
			// BUG crrPth can not update when C:
			//
			// println(tmpDir)
			// if strings.ContainsAny(tmpDir, "\\") == false || strings.ContainsAny(tmpDir, "/") == false {
			// 	crrPth = "C:"
			// } else {
			// 	crrPth = tmpDir
			// }

			crrPth = tmpDir

			dataToCli = "\n" + crrPth + "#\n"

			fmt.Printf(crrPth + "#\n")
			c.Write([]byte(dataToCli))

		} else if strings.HasPrefix(dataFromCli, "... ") == true {

			tmpLetter := ""
			tmpDir := ""

			for i, letter := range dataFromCli {

				tmpLetter += string(letter)

				if tmpLetter == "... " {
					tmpDir = dataFromCli[i+1:]
					break
				}

			}

			// tmpDir =  Dir name to switch, from client.

			//
			// tmpDir control
			//

			println("crrPth:", crrPth)

			filesArr := []string{}

			files, err := ioutil.ReadDir(crrPth)
			if err != nil {
				fmt.Println("[ERR] Read directory:", err)
			}

			for _, file := range files {

				if file.IsDir() {

					filesArr = append(filesArr, file.Name())

				}
			}

			tmpLgc := false

			for _, dir := range filesArr {

				if dir == tmpDir {
					tmpLgc = true
				}

			}

			//
			//
			//

			if tmpLgc == true {

				crrPth += "\\" + tmpDir

				dataToCli = "\n" + crrPth + "#\n"

				fmt.Printf(crrPth + "#\n")
				c.Write([]byte(dataToCli))

			} else {

				dataToCli = "\n[ERR]" + tmpDir + "does not exist.\n"

				fmt.Printf("[ERR] " + tmpDir + " does not exist.\n")
				c.Write([]byte(dataToCli))

			}

		} else {
			// Default
			//
			// It must send "any char". In this case, a space.
			c.Write([]byte(" "))
		}

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
