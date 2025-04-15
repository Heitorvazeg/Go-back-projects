package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Servidor rodando!")

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatalln("Erro ao iniciar servidor UDP!", err)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		n, clientAddr, err := conn.ReadFromUDP(buf)

		if err != nil {
			log.Println("Erro ao ler cliente!", err)
			continue
		}

		fmt.Printf("Mensagem recebida de %s: %s\n", clientAddr, string(buf[:n]))

		_, err = conn.WriteToUDP([]byte("Hye Client! "), clientAddr)

		if err != nil {
			log.Println("Erro ao enviar resposta: ", err)
		}
	}
}
