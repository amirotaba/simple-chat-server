package main

import (
	"bufio"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"os"
	"strings"
)

type message struct {
	Name string
	Text string
}

func main() {
	nc := Conn()
	Sub(nc)
	fmt.Println("enter your name: ")
	name := Read()
	for {
		fmt.Println("type message: ")
		text := Read()
		msg := &message{
			Name: name,
			Text: text,
		}
		if msg.Text != "" {
			text := NewMessage(msg)
			Pub(nc, text)
		}
	}

}

func Conn() *nats.Conn {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)
	if nc != nil {
		log.Println("Connected to " + nats.DefaultURL)
	}
	return nc
}

func Pub(nc *nats.Conn, text string) {
	err := nc.Publish("reply", []byte(text))
	if err == nil {
		log.Println("Message published")
	}
}

func Sub(nc *nats.Conn) {
	_, err := nc.Subscribe("send", func(msg *nats.Msg) {
		log.Println(string(msg.Data))
	})
	if err != nil {
		log.Printf("message didn't publish successfully %v", err)
	}
}

func Read() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func NewMessage(msg *message) string {
	return msg.Name + ": " + msg.Text
}
