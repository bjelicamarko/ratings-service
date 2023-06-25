package producer

import (
	"RatingsService/config"
	"RatingsService/models"
)

var ACCOMMODATION_RATING_QUEUE string = config.ReturnConfig().RatingsQueue
var RESERVATION_QUEUE string = config.ReturnConfig().ReservationsQueue

var typeQueuesMap = map[models.MessageType][]string{
	models.ADD_ACCOMMODATION_RATING_INITIATED: {RESERVATION_QUEUE},
	models.ADD_HOST_RATING_INITIATED:          {RESERVATION_QUEUE},
}
