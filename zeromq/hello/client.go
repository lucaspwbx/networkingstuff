package main

import (
	"fmt"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	fmt.Println("COnnecting")
	requester, _ := zmq.NewSocket(zmq.REQ)
	defer requester.Close()
	requester.Connect("tcp://localhost:5555")
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("Hello %d", i)
		fmt.Println("Sending ", msg)
		requester.Send(msg, 0)

		reply, _ := requester.Recv(0)
		fmt.Println("Received ", reply)
	}
}
