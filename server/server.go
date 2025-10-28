package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type ChatServer struct {
	messages []string
}

// Struct to define message type
type Message struct {
	Name    string
	Content string
}

// Function to send message from client to server
func (c *ChatServer) SendMessage(msg Message, reply *[]string) error {
	fullMsg := fmt.Sprintf("%s: %s", msg.Name, msg.Content)
	c.messages = append(c.messages, fullMsg)
	*reply = c.messages
	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server running on port 1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
