package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("vim-go")
	// var address string
	// flag.StringVar(&address, "address", "127.0.0.1:8000", "Host and port number")
	// flag.Parse()

	// fmt.Printf("Using port: %v\n", address)

	conn, err := net.Dial("udp", "127.0.0.1:8000")
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
		time.Sleep(time.Second * 1)
	}
}
