package main

import (
	"RatingsService/db"
	"RatingsService/handlers"
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
		router.MapRoutesAndServe(handler)
	} else {
		log.Printf("Closing the application, could not connect to db!")
	}
}
