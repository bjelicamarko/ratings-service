package consumer

import (
	"RatingsService/helpers"
	"RatingsService/models"
	"RatingsService/rabbitHandlers"
	"RatingsService/repositories"
	"encoding/json"
	"log"
)

// Map MessageType to handler function

var typeHandlerMap = map[models.MessageType]func(*models.Message, *repositories.RatingsRepository){
	models.ADD_ACCOMMODATION_RATING_SUCCEEDED: rabbitHandlers.AddAccommodationRatingSucceededHandler,
	models.ADD_ACCOMMODATION_RATING_FAILED:    rabbitHandlers.AddAccommodationRatingFailedHandler,
}

func InvokeHandler(data []byte, repository *repositories.RatingsRepository) {
	var message models.Message

	err := json.Unmarshal(data, &message)
	if err != nil {
		log.Println("Error occurred on data unmarshal")
		log.Println(err)
		return
	}

	if handler, ok := typeHandlerMap[message.Type]; ok {
		log.Printf("Message with id: [%d] is passed to handler: [%s]", message.ID, helpers.GetFunctionName(handler))
		handler(&message, repository)
		log.Printf("Message with id: [%d] is successfully handled by handler: [%s]", message.ID, helpers.GetFunctionName(handler))
	} else {
		log.Printf("Handler for message with type: %s does not exist id: [%d]", message.Type, message.ID)
	}
}
