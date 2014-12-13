package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	responder, _ := zmq.NewSocket(zmq.REP)
	defer responder.Close()

	responder.Bind("tcp://*:5555")
	for {
		msg, _ := responder.Recv(0)
		fmt.Println("Received ", msg)

		time.Sleep(time.Second)

		reply := "World"
		responder.Send(reply, 0)
	}
}
