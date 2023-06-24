package rabbitHandlers

import (
	"RatingsService/models"
	"RatingsService/repositories"
	"encoding/json"
	"log"
)

type AddAccommodationRatingFailed struct {
	RatingId uint   `json:"RatingId"`
	Reason   string `json:"reason"`
}

func AddAccommodationRatingFailedHandler(message *models.Message, repository *repositories.RatingsRepository) {
	var body *AddAccommodationRatingFailed
	body, err := unmarshalSucceededForFailedAdd(message, body)

	if err != nil {
		handleSucceededErrorForFailedAdd(message.ID, err)
		return
	}

	rating, err := repository.GetAccommodationRatingById(body.RatingId)
	if err != nil {
		handleSucceededErrorForFailedAdd(message.ID, err)
		return
	}

	rating.Status = models.REJECTED

	err = repository.Update(rating)
	if err != nil {
		handleSucceededErrorForFailedAdd(message.ID, err)
		return
	}

	log.Printf(body.Reason)
	// Notify user through web socket
}

func unmarshalSucceededForFailedAdd(message *models.Message, extractReference *AddAccommodationRatingFailed) (*AddAccommodationRatingFailed, error) {
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

func handleSucceededErrorForFailedAdd(messageId uint, err error) {
	log.Printf("Error occurred when handling ADD_ACCOMMODATION_RATING_FAILED for message with id: [%d]. Error: %s",
		messageId, err.Error())
	// Notify user through web socket
}
