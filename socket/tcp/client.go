package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	scanner := bufio.NewScanner(os.Stdin)

	count := 0

	for {
		fmt.Print("> ")

		if scanner.Scan() {
			msg := scanner.Text() + "\n"
			writer.WriteString(msg)
			writer.Flush()

			response, err := reader.ReadString('\n')

			if err != nil {
				if strings.Contains(err.Error(), "wsarecv") {
					fmt.Println("O servidor encerrou a conexão!")
					break

				} else if err.Error() == "EOF" {
					fmt.Println("Servidor desconectado!")
					break

				} else {
					fmt.Println("Falha ao receber mensagem...", err)

					count++
					if count == 3 {
						fmt.Println("Conexão encerrada por sequência de falhas!")
						break
					}
				}
			} else {
				count = 0
				fmt.Printf("%s", response)
			}

		} else {
			break
		}
	}
}
