// Copyright (C) fothoncode.
// Author: I Made Edy Gunawan
package golang_rabbitmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Config struct {
	RabbitMQString string
	QueueKey       string
}

func CreateQueue(rabbitMQChanel *amqp.Channel, queueKey string) amqp.Queue {
	q, err := rabbitMQChanel.QueueDeclare(
		queueKey, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Queue created %v\n", q)

	return q
}

func Connect(rabbitMQString string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(rabbitMQString)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Connected to %v\n", rabbitMQString)

	return conn, ch
}

func Consume(config Config, handle func(data amqp.Delivery)) {
	mqConn, mqChannel := Connect(config.RabbitMQString)
	defer mqConn.Close()
	defer mqChannel.Close()

	mqQueue := CreateQueue(mqChannel, config.QueueKey)

	msgs, err := mqChannel.Consume(
		mqQueue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		fmt.Println("Go run!")
		for d := range msgs {
			handle(d)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
