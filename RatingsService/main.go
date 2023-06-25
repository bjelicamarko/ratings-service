package main

import (
	"RatingsService/consumer"
	"RatingsService/db"
	"RatingsService/handlers"
	"RatingsService/producer"
	"RatingsService/repositories"
	"RatingsService/router"
	"RatingsService/services"
	"log"
)

func main() {
	log.Printf("Starting the application.")

	dbConn := db.Init()
	if dbConn != nil {
		repository := repositories.NewRepository(dbConn)
		service := services.NewRatingsService(repository)
		handler := handlers.NewRatingsHandler(service)
		go consumer.StartConsumer(repository)
		go producer.StartProducer(dbConn)
		router.MapRoutesAndServe(handler)
	} else {
		log.Printf("Closing the application, could not connect to db!")
	}
}
