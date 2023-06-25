package consumer

import (
	"RatingsService/config"
	"RatingsService/helpers"
	"RatingsService/repositories"
	"log"

	"github.com/streadway/amqp"
)

func StartConsumer(repository *repositories.RatingsRepository) {
	log.Println("Connecting to RabbitMQ instance...")

	rabbitMqUrl := config.ReturnConfig().RabbitMqUrl
	queueName := config.ReturnConfig().RatingsQueue

	conn, err := amqp.Dial(rabbitMqUrl)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer conn.Close()

	log.Println("Successfully connected to RabbitMQ instance")

	chanel, err := conn.Channel()
	if err != nil {
		log.Println(err)
		panic(err)
	}
	defer chanel.Close()

	log.Printf("Creating '%s'...\n", queueName)

	queue, err := chanel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // autoDelete
		false,     // exclusive
		false,     // noWait
		nil,       // args
	)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	log.Printf("Successfully created '%s'...\n", queueName)
	log.Println(queue)

	log.Println("Starting consumer...")

	msgs, err := chanel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // autoAck
		false,     // exclusive
		false,     // noLocal
		false,     // noWait
		nil,       // args
	)

	if err != nil {
		log.Println("Error occurred at consuming events...")
		log.Println(err)
	}

	log.Println("Consumer successfully started...")

	neverend := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Println("Consumed message")
			helpers.LogMessageBody(msg.Body)
			InvokeHandler(msg.Body, repository)
		}
	}()
	<-neverend
}
