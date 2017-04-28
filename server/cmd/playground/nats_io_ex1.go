package main

import (
	"time"
	"fmt"
	"github.com/nats-io/go-nats"
)

func main() {

	nc, error1 := nats.Connect(nats.DefaultURL)
	if error1 != nil {
		fmt.Printf("nats.Connect: %v\n", error1)
	}

	c, error2 := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if error2 != nil {
		fmt.Printf("nats.NewEncodedConn: %v\n", error2)
	}

	defer c.Close()

	// Simple Publisher
	c.Publish("foo", "Hello World")

	// Simple Async Subscriber
	c.Subscribe("foo", func(s string) {
		fmt.Printf("Received a message: %s\n", s)
	})

	// EncodedConn can Publish any raw Go type using the registered Encoder
	type person struct {
		Name     string
		Address  string
		Age      int
	}

	c.Publish("foo", "Hello World1")
	c.Publish("foo", "Hello World2")
	c.Publish("foo", "Hello World3")

	// Go type Subscriber
	c.Subscribe("hello", func(p *person) {
		fmt.Printf("Received a person: %+v\n", p)
	})

	me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street, San Francisco, CA"}

	// Go type Publisher
	c.Publish("hello", me)

	// Unsubscribe
	sub, err := c.Subscribe("foo", nil)
	if err != nil {
		fmt.Printf("c.Subscribe foo Request failed: %v\n", err)
	}

	sub.Unsubscribe()

	// Requests
	var response string
	errorr := c.Request("help", "help me", &response, 10*time.Millisecond)
	if errorr != nil {
		fmt.Printf("c.Requeshelp Request failed: %v\n", errorr)
	}

	// Replying
	c.Subscribe("help", func(subj, reply string, msg string) {
		c.Publish(reply, "I can help!")
	})

	// Close connection
	c.Close();
}
