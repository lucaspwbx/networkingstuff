package main

import (
	"time"

	zmq "github.com/pebbe/zmq4"
)

func main() {
	publisher, _ := zmq.NewSocket(zmq.PUB)
	defer publisher.Close()
	publisher.Bind("tcp://*:5563")
	for {
		publisher.Send("A", zmq.SNDMORE)
		publisher.Send("We dont want to see this", 0)
		publisher.Send("B", zmq.SNDMORE)
		publisher.Send("we would like to see this", 0)
		time.Sleep(time.Second)
	}
}
