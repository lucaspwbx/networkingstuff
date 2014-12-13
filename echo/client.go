package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", ":1200")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = conn.Write([]byte("anything"))
	if err != nil {
		log.Fatalln(err)
	}

	var buf []byte
	buf = make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf[0:n]))
}
