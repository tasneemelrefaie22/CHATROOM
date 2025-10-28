package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Name    string
	Content string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You've joined the chat. Type a message to see the chat history.\n", name)

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		msg := Message{Name: name, Content: text}
		var reply []string
		err = client.Call("ChatServer.SendMessage", msg, &reply)
		if err != nil {
			fmt.Println("Error calling RPC:", err)
			continue
		}

		fmt.Println("\n--- Chat History ---")
		for _, m := range reply {
			fmt.Println(m)
		}
		fmt.Println("--------------------")
	}
}
