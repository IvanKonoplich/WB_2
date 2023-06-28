package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("starting server")
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
			return
		}
	}(conn)

	buf := make([]byte, 1)

	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		_, err = conn.Write(buf)
		if err != nil {
			log.Println(err)
			break
		}
	}
}
