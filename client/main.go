package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Print("Starting client...\n")
	var serverAddress string
	var port string

	flag.StringVar(&serverAddress, "address", "127.0.0.1:8000", "Host and port number")
	flag.StringVar(&port, "port", "8001", "Client port number")
	flag.Parse()

	port = ":" + port

	localAddr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}
	serverAddr, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	i := 0
	for {
		msg := strconv.Itoa(i)
		i++
		buff := []byte(msg)
		_, err := conn.Write(buff)
		if err != nil {
			fmt.Print(err)
		}
		time.Sleep(time.Second * 2)
	}
}
