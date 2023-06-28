package rabbitHandlers

import (
	"RatingsService/models"
	"RatingsService/repositories"
	"encoding/json"
	"log"
)

type AddAccommodationRatingAccepted struct {
	GuestId         uint `json:"GuestId"`
	AccommodationId uint `json:"AccommodationId"`
	RatingId        uint `json:"RatingId"`
}

func AddAccommodationRatingSucceededHandler(message *models.Message, repository *repositories.RatingsRepository) {
	var body *AddAccommodationRatingAccepted
	body, err := unmarshalSucceeded(message, body)

	if err != nil {
		handleSucceededError(message.ID, err)
		return
	}

	rating, err := repository.GetAccommodationRatingById(body.RatingId)
	if err != nil {
		handleSucceededError(message.ID, err)
		return
	}

	rating.Status = models.ACCEPTED

	_, err = repository.UpdateAccommodationRating(rating)
	if err != nil {
		handleSucceededError(message.ID, err)
		return
	}

	err = repository.CreateNotificationAccommodationRated(rating)
	if err != nil {
		return
	}
	log.Printf("Successfully added accommodation rating with id: [%d] for accommodation: [%d] and guest: [%d]", rating.ID, rating.AccommodationId, rating.GuestId)
	// Notify user through web socket
}

func unmarshalSucceeded(message *models.Message, extractReference *AddAccommodationRatingAccepted) (*AddAccommodationRatingAccepted, error) {
	jsonBytes, err := json.Marshal(message.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &extractReference)
	if err != nil {
		return nil, err
	}

	return extractReference, nil
}

func handleSucceededError(messageId uint, err error) {
	log.Printf("Error occurred when handling ADD_ACCOMMODATION_RATING_INITIATED for message with id: [%d]. Error: %s",
		messageId, err.Error())
	// Notify user through web socket
}
