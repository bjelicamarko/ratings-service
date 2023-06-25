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
		GuestId:         1,
		AccommodationId: 1,
	},
	{
		Model:           gorm.Model{},
		Mark:            4,
		Comment:         "odlicno",
		GuestId:         1,
		AccommodationId: 1,
	},
}

var HostRatings = []models.HostRating{
	{
		Model:   gorm.Model{},
		Mark:    5,
		Comment: "onako",
		GuestId: 1,
		HostId:  1,
	},
	{
		Model:   gorm.Model{},
		Mark:    4,
		Comment: "odlicno",
		GuestId: 1,
		HostId:  1,
	},
}
