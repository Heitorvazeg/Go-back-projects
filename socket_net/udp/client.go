package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8081")

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	msg := "Hello Servidor UDP!"

	_, err = conn.Write([]byte(msg))

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Enviado: %s\n", string(msg))

	conn.SetReadDeadline(time.Now().Add(5 * time.Second))

	response := make([]byte, 1024)
	n, err := conn.Read(response)

	if err != nil {
		log.Println("Erro ao ler do servidor: ", err)
		return
	}

	fmt.Println("Recebido: ", string(response[:n]))
}
