package main

import (
	"fmt"
	"testing"
)

type Server struct {
	users map[string]string
}

func NewServer() *Server {
	return &Server{
		users: make(map[string]string),
	}
}

func main() {

	userChan := make(chan string, 2)

	userChan <- "John"
	userChan <- "Mary"
	//userChan <- "FooBar" // here will produce dead lock, because the channel is full.

	user := <-userChan

	fmt.Println(user)

	msgChan := make(chan string)

	sendMessage(msgChan)

	consumeMessage(msgChan)

}

func sendMessage(msgChan chan<- string) {
	msgChan <- "Hello!"
}

func consumeMessage(msgChan <-chan string) {
	msg := <-msgChan
	fmt.Println(msg)
}

func TestAddUser(t *testing.T) {
	server := NewServer()

}
