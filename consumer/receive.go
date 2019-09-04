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

	msg, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to register consumer, %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msg {
			log.Printf("Received message: %s", d.Body)
		}
	}()
	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
