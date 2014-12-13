package main

type Client struct {
	Nick      string
	Incoming  chan string
	Outcoming chan string
}

func NewClient(nick string) *Client {
	return &Client{
		Nick:      nick,
		Incoming:  make(chan string),
		Outcoming: make(chan string),
	}
}

func main() {

}
