package main

import (
	"io"
	"log"
	"net"
)

func main(){
	lis,err:=net.Listen("tcp",":8080")
	defer lis.Close()
	if err!=nil{
		log.Fatal(err)
	} else{
		for {
			conn, err:=lis.Accept()
			if err!=nil{
				log.Fatal(err)
			} else{
				io.WriteString(conn,"Hello Client!! I am server")
			}
			conn.Close()
		}
	}
}