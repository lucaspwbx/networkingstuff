package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", ":1200")
	if err != nil {
		log.Fatalln(err)
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		var buf []byte
		buf = make([]byte, 512)
		for {
			n, err := conn.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(buf[0:n]))
			_, err2 := conn.Write(buf[0:n])
			if err2 != nil {
				log.Fatalln(err2)
			}
		}
		conn.Close()
	}
}
