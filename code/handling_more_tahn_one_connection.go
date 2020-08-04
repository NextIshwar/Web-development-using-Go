package main

import (
	"bufio"
	"fmt"
	_ "io"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	defer lis.Close()
	if err != nil {
		log.Panic(err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("I heard you say", line)
	}
	defer conn.Close()
}
