package db

import (
	"RatingsService/models"

	"gorm.io/gorm"
)

var Ratings = []models.AccommodationRating{
	{
		Model:           gorm.Model{},
		Mark:            5,
		Comment:         "onako",
		GuestId:         2,
		HostId:          1,
		AccommodationId: 1,
		Status:          models.ACCEPTED,
	},
	{
		Model:           gorm.Model{},
		Mark:            4,
		Comment:         "odlicno",
		GuestId:         3,
		HostId:          1,
		AccommodationId: 1,
		Status:          models.ACCEPTED,
	},
}

var HostRatings = []models.HostRating{
	{
		Model:   gorm.Model{},
		Mark:    5,
		Comment: "onako",
		GuestId: 4,
		HostId:  1,
		Status:  models.ACCEPTED,
	},
	{
		Model:   gorm.Model{},
		Mark:    4,
		Comment: "odlicno",
		GuestId: 5,
		HostId:  1,
		Status:  models.ACCEPTED,
	},
}
