package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Print("Starting server...\n")
	var serverAddress string

	flag.StringVar(&serverAddress, "address", "127.0.0.1:8000", "Host and port to listen on")
	flag.Parse()

	packetConn, err := net.ListenPacket("udp", serverAddress)

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
