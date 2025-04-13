package main

import (
	"log"
	"net"
)

func handleConnection() {

}

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8081")

	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		n, clientAddr, err := conn.ReadFromUDP(buf)

		if err != nil {
			log.Fatalln(err)
		}

		_, err = conn.WriteToUDP([]byte("Hye Client!", clientAddr))

		if err != nil {
			log.Println("Error writing: ", err)
		}
	}
}
