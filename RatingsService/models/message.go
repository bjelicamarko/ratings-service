package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Type      MessageType    `gorm:"not null"`
	Published bool           `gorm:"not null"`
	Body      datatypes.JSON `gorm:"type:jsonb; not null"`
}

// ********** MessageType Definitions ********** //

type MessageType string

const (
	ADD_ACCOMMODATION_RATING_INITIATED MessageType = "ADD_ACCOMMODATION_RATING_INITIATED"
	ADD_ACCOMMODATION_RATING_SUCCEEDED MessageType = "ADD_ACCOMMODATION_RATING_SUCCEEDED"
	ADD_ACCOMMODATION_RATING_FAILED    MessageType = "ADD_ACCOMMODATION_RATING_FAILED"

	ADD_HOST_RATING_INITIATED MessageType = "ADD_HOST_RATING_INITIATED"
	ADD_HOST_RATING_SUCCEEDED MessageType = "ADD_HOST_RATING_SUCCEEDED"
	ADD_HOST_RATING_FAILED    MessageType = "ADD_HOST_RATING_FAILED"
)

// ********** MessageBodyTypesForPublishing ********** //

type AddAccommodationRatingInitiated struct {
	GuestId         uint `json:"GuestId"`
	AccommodationId uint `json:"AccommodationId"`
	RatingId        uint `json:"RatingId"`
}

type AddHostRatingInitiated struct {
	GuestId  uint `json:"GuestId"`
	HostId   uint `json:"HostId"`
	RatingId uint `json:"RatingId"`
}
