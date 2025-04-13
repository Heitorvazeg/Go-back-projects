package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Servidor rodando!")

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		fmt.Println("Conexão aceita!")

		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		data, err := reader.ReadString('\n')

		if err != nil {
			log.Fatalln(err)
			return

		} else if data == "Vou sair!\n" {
			fmt.Println("Cliente encerrou conexão")
			break
		}

		response := fmt.Sprintf("Echo: %s", data)
		writer.WriteString(response)
		writer.Flush()
	}
}
