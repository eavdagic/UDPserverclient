package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("vim-go")

	packetConn, err := net.ListenPacket("udp", "127.0.0.1:8000")

	if err != nil {
		log.Fatal(err)
	}
	defer packetConn.Close()

	buffer := make([]byte, 1024)

	for {
		n, addr, err := packetConn.ReadFrom(buffer)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println("Received ", string(buffer[0:n]), " from ", addr)
	}
}
