package rabbitHandlers

import (
	"RatingsService/models"
	"RatingsService/repositories"
	"encoding/json"
	"log"
)

type AddHostRatingAccepted struct {
	GuestId  uint `json:"GuestId"`
	HostId   uint `json:"HostId"`
	RatingId uint `json:"RatingId"`
}

func AddHostRatingSucceededHandler(message *models.Message, repository *repositories.RatingsRepository) {
	var body *AddHostRatingAccepted
	body, err := unmarshalSucceededAddHostRating(message, body)

	if err != nil {
		handleSucceededErrorAddHostRating(message.ID, err)
		return
	}

	rating, err := repository.GetHostRatingById(body.RatingId)
	if err != nil {
		handleSucceededErrorAddHostRating(message.ID, err)
		return
	}

	rating.Status = models.ACCEPTED

	_, err = repository.UpdateHostRating(rating)
	if err != nil {
		handleSucceededErrorAddHostRating(message.ID, err)
		return
	}

	log.Printf("Successfully added host rating with id: [%d] for host: [%d] and guest: [%d]", rating.ID, rating.HostId, rating.GuestId)
	// Notify user through web socket
}

func unmarshalSucceededAddHostRating(message *models.Message, extractReference *AddHostRatingAccepted) (*AddHostRatingAccepted, error) {
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

func handleSucceededErrorAddHostRating(messageId uint, err error) {
	log.Printf("Error occurred when handling ADD_HOST_RATING_INITIATED for message with id: [%d]. Error: %s",
		messageId, err.Error())
	// Notify user through web socket
}
