package main

import (
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		log.Fatalf("Failed to connect to rabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("anvh", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue, %v", err)
	}

	body := "Hello World"
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		log.Fatalf("Failed to publish message, %v", err)
	}
}
