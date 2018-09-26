package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"hash/adler32"
	"log"
	"net"
	"time"
)

type Payload struct {
	Msg  string
	Ip   string
	Port string
	Time time.Time
}

func listen(conn *net.UDPConn, buf []byte) {
	n, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Print(err)
	}

	buffer := bytes.NewBuffer(buf[0:n])
	decoder := gob.NewDecoder(buffer)

	var pl Payload
	decoder.Decode(&pl)

	_ = adler32.Checksum(buf)
	fmt.Printf("Server received quote: %s\n", pl.Msg)
	fmt.Printf("Server received source address: %s\n", pl.Ip)
}

func main() {
	fmt.Print("Starting server...\n")
	var serverAddress string

	flag.StringVar(&serverAddress, "address", "127.0.0.1:8000", "Host and port to listen on")
	flag.Parse()

	udpAddress, err := net.ResolveUDPAddr("udp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	ln, err := net.ListenUDP("udp", udpAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	buf := make([]byte, 1024)

	for {
		listen(ln, buf)
	}
}
