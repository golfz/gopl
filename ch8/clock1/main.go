package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	fmt.Println("handleConn")
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		fmt.Println(time.Now())
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
