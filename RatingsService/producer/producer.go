package producer

import (
	"RatingsService/config"
	"RatingsService/helpers"
	"RatingsService/models"
	"encoding/json"
	"log"
	"time"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func StartProducer(db *gorm.DB) {
	rabbitMqUrl := config.ReturnConfig().RabbitMqUrl

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

	for {
		publish(db, chanel)
	}
}

func publish(db *gorm.DB, chanel *amqp.Channel) {
	db.Transaction(func(tx *gorm.DB) error {
		var messages []models.Message

		err := tx.Select("*").
			Table("messages").
			Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("published = ?", false).
			Find(&messages).
			Limit(5).
			Error

		if err != nil {
			log.Println("Error occurred while reading messages from database")
			log.Println(err)
			return err
		}

		if len(messages) == 0 {
			time.Sleep(2 * time.Second)
			return nil
		}

		for _, message := range messages {
			queueNames, ok := typeQueuesMap[message.Type]

			if !ok {
				log.Printf("Array of queues for message with type: %s is undefined id: [%d]", message.Type, message.ID)
				continue
			}

			if len(queueNames) == 0 {
				log.Printf("Array of queues for message with type: %s is empty id: [%d]", message.Type, message.ID)
				continue
			}

			marshall, err := json.Marshal(message)
			if err != nil {
				log.Printf("Could not marshall message with id: [%d]", message.ID)
				return err
			}

			for _, queueName := range queueNames {
				err = chanel.Publish(
					"",
					queueName,
					false,
					false,
					amqp.Publishing{
						Body: marshall,
					},
				)

				if err != nil {
					log.Printf("Failed to publish message with id: [%d] to %s", message.ID, queueName)
					log.Println(err)
					return err
				}

				log.Printf("Successfully published message with id: [%d] to [%s]", message.ID, queueName)
				helpers.LogMessageBody(marshall)
			}

			message.Published = true
			tx.Model(&message).Update("published", true)
		}
		return nil
	})

}
