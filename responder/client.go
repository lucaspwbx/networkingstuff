package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":1200")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}
	write(conn)
	read(conn)
	write(conn)
	read(conn)
	//read(conn)
}

func write(conn net.Conn) {
	conn.Write([]byte("Inter"))
}

func read(conn net.Conn) {
	var buf = make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatalln(err)
	}
	fmt.Println(string(buf[:n]))
}
