package rabbitHandlers

import (
	"RatingsService/models"
	"RatingsService/repositories"
	"encoding/json"
	"log"
)

type AddHostRatingFailed struct {
	RatingId uint   `json:"RatingId"`
	Reason   string `json:"reason"`
}

func AddHostRatingFailedHandler(message *models.Message, repository *repositories.RatingsRepository) {
	var body *AddHostRatingFailed
	body, err := unmarshalSucceededForFailedAddHostRating(message, body)

	if err != nil {
		handleSucceededErrorForFailedAddHostRating(message.ID, err)
		return
	}

	rating, err := repository.GetHostRatingById(body.RatingId)
	if err != nil {
		handleSucceededErrorForFailedAddHostRating(message.ID, err)
		return
	}

	rating.Status = models.REJECTED

	_, err = repository.UpdateHostRating(rating)
	if err != nil {
		handleSucceededErrorForFailedAddHostRating(message.ID, err)
		return
	}

	log.Printf(body.Reason)
	// Notify user through web socket
}

func unmarshalSucceededForFailedAddHostRating(message *models.Message, extractReference *AddHostRatingFailed) (*AddHostRatingFailed, error) {
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

func handleSucceededErrorForFailedAddHostRating(messageId uint, err error) {
	log.Printf("Error occurred when handling ADD_HOST_RATING_FAILED for message with id: [%d]. Error: %s",
		messageId, err.Error())
	// Notify user through web socket
}
