package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
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

func newPayload(msg string, port string, ip string) Payload {
	currentTime := time.Now()
	pl := Payload{
		Msg:  msg,
		Time: currentTime,
		Port: port,
		Ip:   ip,
	}

	return pl
}

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

	msg := `"It always seems impossible until it's done" - Nelson Mandela`

	i := 1
	for {
		var buffer bytes.Buffer
		encoder := gob.NewEncoder(&buffer)

		pl := newPayload(msg, port, serverAddress)
		encoder.Encode(pl)

		conn.Write(buffer.Bytes())
		if err != nil {
			fmt.Print(err)
		}
		msgSentAt := time.Now()

		time.Sleep(time.Second * 2)
		fmt.Print("------------------\n")
		fmt.Printf("Time elapsed since previous message was sent: %s\n", time.Since(msgSentAt))
		fmt.Printf("Total messages sent so far: %d\n", i)
		i++
	}
}
