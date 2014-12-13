package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var answers = map[string]string{
	"Inter":  "Campeao de Tudo",
	"Gremio": "Campeao de Nada",
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp", ":1200")
	if err != nil {
		log.Fatalln(err)
	}
	//listen
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Println(err)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		for {
			read(conn)
		}
	}
}

func read(conn net.Conn) {
	var buf = make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Read", string(buf[:n]))
	if answer, ok := answers[string(buf[:n])]; ok {
		write(conn, answer)
	}
}

func write(conn net.Conn, answer string) {
	conn.Write([]byte("We've got an answer for you: " + answer))
}
